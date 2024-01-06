# User-Event-Aggregation-CLI
This command-line utility aggregates user activity events from a JSON file and generates daily summary reports. It also supports real-time updates as new events are added to the dataset.

## Project Structure
- aggregate_events.go:
   - The main Go source code file containing the implementation of the user event aggregation utility.
- aggregate_events_test.go: 
   - The test file containing test cases for the user event aggregation utility.
- fileprocessor:
   - A package providing utility functions for reading and writing JSON files.


## Usage

1: Clone the repository:
- git clone https://github.com/himanshu07070/User-Event-Aggregation-CLI

2: Navigate to the project directory:
- cd User-Event-Aggregation-CLI
- Build the Project:
  - go build
- Run the Project(new events or updated events):
  - ./{name_of_the_build_file}
