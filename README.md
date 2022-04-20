# Allure Tool

Utility to extract info from Allure reports. 

With Allure, you can mark tests in your test suite with labels for epic/feature/story, so you can keep track of the existing tests related to user stories in Jira. With this utility, you can search for specific labels and get a report with the tests that are related to them.

```
go build allure-tool.go
```

# Usage

By default, the tool expects the following structure:

```
.
├── README.md
├── allure-tool
├── data
│   ├── allure
│   │   ├── suite-behaviors-latest.csv
│   │   ├── other-suite-behaviors-latest.csv
│   │   └── more-tests.csv
│   └── filters.csv
├── go.mod
```

The filters.csv is a simple file the labels you want to extract:

```csv
US-123
US-3213
US-437
```

If you have this structure and execute the tool:

```
./allure-tool
```

You will get an **output.csv** file in the data directory. The tool will process all the files, but it will generate an unified report.

You can pass the following options to customize your environment.

Change base folder for your files instead of `data`.

```
./allure-tool -base otherFolder
```

Change the name of the folder that contains your reports, instead of `allure`:

```
./allure-tool -source reports
```

Use another filters file:

```
./allure-tool -filters alternative-filters.csv
```

Or change the name of the output file:

```
./allure-tool -output filtered-stories.csv
```