# 1password to Enpass Migration

This is a utility helping with migration from 1password to Enpass.

It converts 1password CSV export file to Enpass Pre-formatted CSV file with some normalization.

## Usage

1. Export 1password data to CSV file.
1. Run the utility.
1. Import Pre-formatted CSV file in Enpass.
1. Enjoy normalized data in Enpass without further manual modifications.

*Warning: Backup your original files in case you'll be missing some extra fields not included in a new file.*

## Requirements

Input file (exported from 1password) has to be named `input.csv` and its path
must be `./input.csv` (in current working directory when launching the utility).

## Notes

This is just a one-time use utility and it does not contain any tests, code quality checks, CI, etc.

I don't want to waste time making this super duper rocket science app.

Feel free to contribute or just fork it and modify to your needs.