
## Requirements
The microservice logs each API request to a text log file, stored at:

> /var/log/httpd/*.log 

The log files are rotated every day at midnight, so there are multiple days of logs in that directory.

Each log file contains multiple log lines, and the format for each log line is:

    IP_ADDRESS [timestamp] "HTTP_VERB URI" HTTP_ERROR_CODE RESPONSE_TIME_IN_MILLISECONDS

## Code structure
The project can be composed of by three parts
- parse log fields and add values to sorter
- get percentile from added values (use _github.com/influxdata/tdigest_ implementation)
- print results

a parameter can use to custom compression which to control trade-off between the size of t-digest and quantiles accuracy.

    td := tdigest.NewWithCompression(1000)

## Algorithm
According to [elastic blog](https://www.elastic.co/blog/averages-can-dangerous-use-percentile "Averages Can Be Misleading: Try a Percentile")
> it is sufficient to make the following claims about T-Digest:
>  
> - For small datasets, your percentiles will be highly accurate (potentially 100% exact if the data is small enough)
> - For larger datasets, T-Digest will begin to trade accuracy for memory savings so that your node doesn't explode
> - Extreme percentiles (e.g. 95th) tend to be more accurate than interior percentiles (e.g. 50th)

[T-Digest Paper](https://github.com/tdunning/t-digest/blob/master/docs/t-digest-paper/histo.pdf)

### Time complexity and space complexity
> If c<sub>1</sub> is the input buﬀer size, the dominant costs are the sort and the scale function calls so the amortized cost per input value is roughly C<sub>1</sub>logc<sub>1</sub>+ C<sub>2</sub>⌈δ⌉/c<sub>1</sub> where C<sub>1</sub> and C<sub>2</sub> are parameters representing the sort and scale function costs respectively.

## Testing
    make test

### Algorithm testing
Add values from 1 to 10000

    go test pkg/tdigest/influxdata_test.go -v

Testing Result

    === RUN   TestInfluxdata
    2020/06/23 16:04:29 50th 5000.5
    2020/06/23 16:04:29 75th 7500.5
    2020/06/23 16:04:29 90th 9000.5
    2020/06/23 16:04:29 99th 9900.5
    --- PASS: TestInfluxdata (0.00s)

### Log parser testing
    go test pkg/file_reader_test.go pkg/file_reader.go

## User guide
1- [Prepare go environment](https://golang.org/doc/install)

2- Extract Log_Percentiles-master.zip file

3- Build binary

    make build
4- Run

    ./bin/main -h
    Usage of ./bin/main:
      -dir string
        	log directory (default ".")

Specify log directory with sample files

    ./bin/main -dir=log