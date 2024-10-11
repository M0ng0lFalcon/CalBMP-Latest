# CalBMP backend

### 整理思路

#### step 1 选择 县郡 跟 zipcode

1. 利用 `zip_county`表
2. `PO_NAME` :  县郡的名字，先选择这个
3. `ZIP_CODE` : 第二步选择这个，一个县郡可能对应多个 zipcode

#### step 2 选择气象站

1. 利用 `zipcode_station` 表
2. 通过 `zip_code` 获取 `climateid` , `log`, `lat` 等数据

#### step 3 选择年份

1. 利用 `data_?` 表
2. 这些表有 `climateid  zipcode  log lat year` 还有一些其他数据

### 参数格式

1. 用什么格式

   目前的方法： 一个一个传params

   新方法：利用json传数据比较好，这样取数据比较方便，然后可以通过一个方法存到后台，可以比较方便管理起来。对了，可以用redis，但是先不用那个了。先用文件来存下来吧。

