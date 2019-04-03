# Trying out Go + Benchmark with PHP

First steps in Go. Comparing results with a similar PHP script.

Converts a supported image into base64 data uri string.

## Benchmark

[hyperfine](https://github.com/sharkdp/hyperfine) was used for the benchmarks.

| PHP (PHP 7.3.3)  | Go (go1.12.1 darwin/amd64) | Node (v11.13.0) |
| ------------- | ------------- | ------------- |
| User: 20.6ms, System: 9.8ms - 74 runs | User: 5.1ms, System: 2.2ms - 286 runs | User: 72.9ms, System: 18.7ms - 30 runs |
|<img width="1122" alt="Screenshot 2019-04-03 at 02 24 04" src="https://user-images.githubusercontent.com/6123841/55442537-becd1c00-55b7-11e9-8d53-5ab7d3f207ad.png">|<img width="1122" alt="Screenshot 2019-04-03 at 02 25 08" src="https://user-images.githubusercontent.com/6123841/55442538-becd1c00-55b7-11e9-96c8-3591b8695f36.png">| <img width="1122" alt="Screenshot 2019-04-03 at 03 57 33" src="https://user-images.githubusercontent.com/6123841/55445574-b16a5e80-55c4-11e9-84af-aa0c6bfca89c.png">
 |
