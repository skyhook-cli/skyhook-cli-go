import org.jenkinsci.plugins.pipeline.modeldefinition.Utils

def REPO_URL = "https://github.com/skyhook-cli/skyhook-cli-go.git"

def COMMIT_MESSAGE
def VERSION_NUMBER

node {
    properties([[$class: 'JiraProjectProperty'], buildDiscarder(logRotator(artifactDaysToKeepStr: '', artifactNumToKeepStr: '', daysToKeepStr: '', numToKeepStr: '5')),
                [$class: 'RebuildSettings', autoRebuild: false, rebuildDisabled: false]])

    deleteDir()

    stage("PR TITLE CHECK") {
        if (isPr()) {
            assert env.CHANGE_TITLE ==~ /(patch|minor|major):.+/
        } else {
            Utils.markStageSkippedForConditional(STAGE_NAME)
        }
    }

    stage("GIT CHECKOUT") {
        git(
            url: "${REPO_URL}",
            credentialsId: 'git-login',
            branch: isPr() ? env.CHANGE_BRANCH : env.BRANCH_NAME
        )

        COMMIT_MESSAGE = sh(
            script: "git log --format=%B -n 1 HEAD",
            returnStdout: true
        ).trim()

        // VERSION UPDATE CODE
        def pushType = COMMIT_MESSAGE.split()[0].toLowerCase().replace(":", "")

        def currentVersion = sh(
            script: "git describe --abbrev=0",
            returnStdout: true
        ).trim()

        currentVersion.replace("v", "").replace("-release", "").tokenize(".")

        println "Current Version: ${currentVersion}"

        switch(pushType) {
            case "patch":
                currentVersion[2] = currentVersion[2].toInteger() + 1
                break
            case "minor":
                currentVersion[1] = currentVersion[1].toInteger() + 1
                break
            case "major":
                currentVersion[0] = currentVersion[0].toInteger() + 1
                break
            default:
                break
        }

        println "Current Version: ${currentVersion}"

        VERSION_NUMBER = currentVersion.join(".")

        println "New Version: ${VERSION_NUMBER}"
        // END VERSION UPDATE CODE

        sh "exit 1"
    }

    stage("GO BUILD") {
        sh "make bins"
    }

    stage("VERSION UPDATE") {
        if (isPushToMaster()) {
            def pushType = COMMIT_MESSAGE.split()[0].toLowerCase().replace(":", "")

            def currentVersion = sh(
                script: "git describe --abbrev=0",
                returnStdout: true
            ).trim()

            currentVersion = currentVersion.replace("v", "").replace("-release", "").tokenize(".")

            println "Current Version: ${currentVersion}"

            switch(pushType) {
                case "minor":
                    println "Minor update"
                    currentVersion[1] = currentVersion[1].toInteger() + 1
                    break
                case "major":
                    println "Major update"
                    currentVersion[0] = currentVersion[0].toInteger() + 1
                    break
                case "patch":
                default:
                    println "Patch update"
                    currentVersion[2] = currentVersion[2].toInteger() + 1
                    break
            }

            VERSION_NUMBER = currentVersion.join(".")

            println "New Version: ${VERSION_NUMBER}"

            withCredentials([
                usernamePassword(credentialsId: 'git-login', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')
            ]) {
                def origin = "https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/skyhook-cli/skyhook-cli-go.git"
                sh """
                    git config user.name "Skyhook Bot"
                    git config user.email "skyhookbot"

                    git tag -a ${VERSION_NUMBER} -m "New tag v${VERSION_NUMBER}"

                    git push ${origin} ${VERSION_NUMBER}
                """
            }
        } else {
            Utils.markStageSkippedForConditional(STAGE_NAME)
        }
    }

    stage("ZIP BINARIES") {
        if (isPushToMaster()) {
            sh "VERSION_NUMBER=${VERSION_NUMBER} make zips"
        } else {
            Utils.markStageSkippedForConditional(STAGE_NAME)
        }
    }

    stage("GITHUB RELEASE") {
        if (isPushToMaster()) {
            withCredentials([
                usernamePassword(credentialsId: 'git-login', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')
            ]) {
                sh """
                    curl https://api.github.com/repos/skyhook-cli/skyhook-cli-go/releases \
                    -H "Authorization: token ${GIT_PASSWORD}" \
                    -H "Accept: application/vnd.github.v3+json" \
                    -H "Content-Type: application/json" \
                    -X POST \
                    -d '{
                        "tag_name": "v${VERSION_NUMBER}-release",
                        "target_commitish": "master",
                        "name": "Release v${VERSION_NUMBER}",
                        "body": "Automated release v${VERSION_NUMBER}"
                    }'
                """
            }

            def id = getReleaseId()

            publishArtifacts(id, "windows", VERSION_NUMBER)
            publishArtifacts(id, "linux", VERSION_NUMBER)
            publishArtifacts(id, "macos", VERSION_NUMBER)
        } else {
            Utils.markStageSkippedForConditional(STAGE_NAME)
        }
    }

}

def isPr() {
    return env.BRANCH_NAME.startsWith("PR-")
}

def isPushToMaster() {
    return env.BRANCH_NAME == "master"
}

def getReleaseId() {
    withCredentials([
        usernamePassword(credentialsId: 'git-login', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')
    ]) {
        return sh(
            script: """
                response=\$(curl -s https://api.github.com/repos/skyhook-cli/skyhook-cli-go/releases/latest -H "Authorization: token ${GIT_PASSWORD}")

                echo \$response | jq .id
            """,
            returnStdout: true
        ).trim()
    }
}

def publishArtifacts(id, os, version) {

    filename = "skyhook-cli-go-${os}-v${version}-x64.zip"
    fullPath = "bin/${os}/${filename}"

    print "Publishing ${filename} for release ${id}"

    withCredentials([
        usernamePassword(credentialsId: 'git-login', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')
    ]) {
        sh """
            curl "https://uploads.github.com/repos/skyhook-cli/skyhook-cli-go/releases/${id}/assets?name=${filename}" \
            -H "Accept: application/vnd.github.v3+json" \
            -H "Content-Type: application/zip" \
            -H "Authorization: token ${GIT_PASSWORD}" \
            -X POST \
            --data-binary @"${fullPath}"
        """
    }
}
