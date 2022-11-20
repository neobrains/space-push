# SpacePush

```bash
name: Push to Space
on: push

jobs:
  push:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: jnsougata/space-push@v0.0.2
        with:
          space-access-token: ${{ secrets.SPACE_ACCESS_TOKEN }}
          space-project-id: ${{ secrets.SPACE_PROJECT_ID }}
```
