<p align="center">
  <p align="center">
    <img src="media/logo.png" width="350" />
  </p>
  
  <p align="center">
    <b>Python Check Updates</b>
  </p>

  <p align="center"></p>
  
  <p align="center">
    <img src="https://img.shields.io/github/go-mod/go-version/xjh22222228/python-check-updates" />
    <img src="https://img.shields.io/github/v/release/xjh22222228/python-check-updates" />
    <img src="https://img.shields.io/github/license/xjh22222228/python-check-updates" />
  </p>
</p>


## PCU
Find the latest version of your requirements.txt current dependency package.


![](media/screenshot.png)






## Usage

Show any new dependencies for the project in the current directory:

```bash
$ pcu
Checking /opt/requirements_test.txt
11 / 11 [------------] 100.00% 1 p/s
pytest            →     6.1.2
pytest-cov        →    2.10.1
pytest-httpbin    →     1.0.0
pytest-mock       →     3.3.1
httpbin           →     0.7.0
wheel             →    0.35.1
alembic           →     1.4.3
appnope           →     0.1.0
astroid           →     2.4.2
attrs             →    20.3.0
backcall          →     0.2.0

Done in  3 s.
```


## Options

```
-f, --file                   Specify the file name of the check dependency
 package, default 'requirements.txt'
-v, --version                output the version number
-h, --help                   display help for command
```

## Example
```
$ pcu -f=requirements_test.txt

$ pcu -v
```






## License
[MIT](LICENSE)

