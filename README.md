# Space Push: Github Action

## Usage
Continuous delivery of [Deta Space Apps](https://alpha.deta.space/)

## Inputs
### `space-access-token`
Used for CLI auth. From [Deta Space Dashboard](https://alpha.deta.space), open the teletype (command bar) and click on `Settings`. In the settings modal, click on `Generate Token` to generate an access token and copy the resulting token.
### `space-project-id`
Used to link current repo with space app. To get your project id go to the [Builder](https://alpha.deta.space/builder) and in your project’s `Develop` tab, open the teletype (command bar) and copy the Project ID.

<br/><br/>
⚠️ Use [GitHub Actions Secrets](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/creating-and-using-encrypted-secrets) to store your inputs ⚠️

## Sample Workflow
`push.yml`
```bash
name: Push to Space
on: push

jobs:
  push:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: neobrains/space-push@v0.0.1
        with:
          space-access-token: ${{ secrets.SPACE_ACCESS_TOKEN }}
          space-project-id: ${{ secrets.SPACE_PROJECT_ID }}
```

Release needs human interaction to be created, thus it is not possible to create a release using GitHub actions at this moment. Follow the instructions [here](https://alpha.deta.space/docs/en/basics/releases#releasing-from-the-gui) to create a release for your project.

##### ** Special thanks to [Sponge](https://github.com/rohanshiva) & [sub-G](https://github.com/mikBighne98)