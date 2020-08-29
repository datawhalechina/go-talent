## 贡献规范

## 0.Readme规范

重点突出：

- 内容
- 目标
- 对象

## 1.提交规范

提交指的是commit规范。

提交规范如下：

- 常规提交

```
git commit -m "normal: 功能 --creator=名字"
```

注意功能与冒号有个空格，功能与`--`有个空格，提交的时候更改功能为自己的提交内容即可，名字替换为自己名字。

例如：光城增加了go if 语句

```
git commit -m "normal: add go if --creator=light-city"
```

- bug修复

```
git commit -m "bugfix: 功能 --creator=名字"
```

将上述的normal替换为bugfix即可。

例如：光城修复了图片不展示问题

```
git commit -m "bugfix: pic show --creator=light-city"
```

## 2.协作规范

指的是pr规范，每个人以pr形式合入master，合入过程中，管理员会对代码或者文档、相关规范进行审查，如果不通过，驳回，否则合入。

每人开发可以自开分支，分支规范如下：

名字/功能

例如：光城分支完成了拉取数据功能。

```
light-city/getData
```

最后，要经常更新master分支，不要让自己分支存活太久。