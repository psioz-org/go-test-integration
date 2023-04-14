# go-test-integration
[![Go Report Card](https://goreportcard.com/badge/github.com/zev-zakaryan/go-test-integration)](https://goreportcard.com/report/github.com/zev-zakaryan/go-test-integration)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/zev-zakaryan/go-test-integration/blob/main/LICENSE)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=zev-zakaryan_go-test-integration&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=zev-zakaryan_go-test-integration)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=zev-zakaryan_go-test-integration&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=zev-zakaryan_go-test-integration)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=zev-zakaryan_go-test-integration&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=zev-zakaryan_go-test-integration)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=zev-zakaryan_go-test-integration&metric=coverage)](https://sonarcloud.io/summary/new_code?id=zev-zakaryan_go-test-integration)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=zev-zakaryan_go-test-integration&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=zev-zakaryan_go-test-integration)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=zev-zakaryan_go-test-integration&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=zev-zakaryan_go-test-integration)

1. Create new project from empty or another project as template (not clone)
2. Change project name related string in README.md, sonar-project.properties and language related files such as go.mod
3. <sup>[1]</sup>Add github variables (as secrets). It can be any level i.e. repo/env/org:
    - Add SONAR_TOKEN: ...from... https://sonarcloud.io/projects/create  
    We use "bound" sonarcloud so we must import project in sonarcloud before (analyze action will fail after copy repo with workflow)  
    Don't create SONAR_TOKEN yourself but get it from auto-created after import project.
        * If you use shared token, make sure user the token was generated for still have Execute Analysis permission on the repos
    - Add REPO_ACTION_TOKEN: ...from... https://github.com/settings/tokens  
    GitHub actions got cannot trigger another GitHub action. Use personal access token so the PRs are opened by yourself thereby triggering the required workflows.  
    https://github.com/googleapis/release-please/issues/922  
    https://github.com/marketplace/actions/release-please-action#github-credentials
        * Need configure in organization to use with github organization repository.
4. Settings
    - Settings > General > "Always suggest updating pull request branches", "Allow auto-merge", "Automatically delete head branches".
    - <sup>[1]</sup>Settings > Actions > General > "Allow GitHub Actions to create and approve pull requests" to use "release-please" action.
    - <sup>[1]</sup>Settings > Code security and analysis > "Private vulnerability reporting"/"Secret scanning": enable.
5. Re-run failed jobs. Wait until completely success
6. Settings > Branches > Add branches rule on main branch as needed:  
    Require status checks:
    - Analyze ({language})[GitHub Actions] *If you enable CodeQL
    - SonarCloudChk[GitHub Actions]
    - GosecChk[GitHub Actions]
    - <sup>[2]</sup> SonarCloud[SonarCloud] *SonarCloud[SonarCloud], gosec[Github Code Scanning] will visible only if start PR
    - gosec[GitHub Code Scanning]
    - <sup>[2]</sup> SonarCloud Code Analysis[GitHub Code Scanning]
7. Merge


<sup>[1]</sup> Can be set only once at organization level  
<sup>[2]</sup> Can omit because it sometimes "Waiting for status to be reported"

Note: GitHub "required workflow" for organization still has many restriction so we still use workflow file in each repository.
