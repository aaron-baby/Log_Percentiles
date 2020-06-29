
## Requirements
The microservice logs each API request to a text log file, stored at:

> /var/log/httpd/*.log 

The log files are rotated every day at midnight, so there are multiple days of logs in that directory.

Each log file contains multiple log lines, and the format for each log line is:

    IP_ADDRESS [timestamp] "HTTP_VERB URI" HTTP_ERROR_CODE RESPONSE_TIME_IN_MILLISECONDS

## Code structure
The project can be composed of by three parts
- parse log fields and add values to sorter
- sort response time
- print results

## Main Process
- sort all log files lines

## Algorithm
According to [elastic blog](https://www.elastic.co/blog/averages-can-dangerous-use-percentile "Averages Can Be Misleading: Try a Percentile")
> it is sufficient to make the following claims about T-Digest:
>  
> - For small datasets, your percentiles will be highly accurate (potentially 100% exact if the data is small enough)
> - For larger datasets, T-Digest will begin to trade accuracy for memory savings so that your node doesn't explode
> - Extreme percentiles (e.g. 95th) tend to be more accurate than interior percentiles (e.g. 50th)
      
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
