# bagoScan

`bagoscan` is a simple tool to run multiple security scans in parallel on a given target.
It currently supports `testssl` and `nmap` scans, but can easily be extended 
to include other tool.
It is written in Go as part of a bachelor's degree project at HTW Berlin.

## Installation

To use `bagoScan`, you'll need to have `Go` installed on your system.
Clone the repository and navigate to the root directory.

```
git clone https://github.com/pomcom/bagoScan.git
cd bagoScan
```
### Docker

You can also use `bagoScan` via Docker by building the image 
with the provided `Dockerfile`

`docker build -t bagoScan .`

## Usage

You can run `bagoScan` using the provided `Makefile` or by running
the binary or via the Docker container.


## Make commands
* `make run` runs the tool
* `make build`build the `bagoScan` binary
* `make clean` removes the binary and the output folder

## Docker commands

You can run `bagoScan` via the built Docker image:

`docker run -it bagoScan -t <target>`
or
`docker run bagoScan bagoscan scan -t <target>`


## Output

The output of the scan will be saved in the `output/raw` directory,
with the filename in the format `TOOLNAME-OUTPUT.txt`.
The output file will also include a timestamp in the format `YYYY-MM-DD-HH-MM-SS`.


## Adding new tools

To add a new tool to `bagoScan`, you'll need to implement the Tool interface in the tools package.
This requires the `Execute(target string) (string, error)` Method to be implemented. 
Once you've done this, you can add the new tool to the `configHandler` in the `utils` package.


## Configuration

You can configure `bagoScan` by providing a configuration file in `yaml` format.
Default path for the `config.yml` is the apps root directory.
Todo: Pass config file via `-c` flag.

Example `config.yml`
```
# define the tools you want to execute
tools:
  - nmap
  - testssl

nmap:
  - -p-
  - -T4

testsll:
  - yourPreferredFlag
  - ...

```



## License

`bagoScan` is released under the [MIT License](https://github.com/pomcom/bagoScan/blob/main/LICENSE).
