# 初始化仓库的时候发生了什么?

## 运行 `git init` 新建仓库

首先我们打开终端, 使用终端复用器分屏一下, 在右边输入

```
watch -d 500 . {|| clear | eza --all --tree --color=always }
```

这将创建一个进程, 每隔 500 ms 读取一次**当前文件夹下的文件** (watch -d 500), 如果有变化, 将会使用 eza 进行输出. 这里 eza 将会输出所有的文件并以树形结构显示.

然后在左边输入 `git init` 初始化一个 git 仓库, 我们将看到以下变化:
![[Pasted image 20250731135551.png]]
 
### `.git/config` 文件

![[Pasted image 20250731104711.png]]

当我们运行 `git config -l` 查看配置时, 会发现配置的 **user.name** 和 **user.email**.

仓库下的  `.git/config` 文件是**仓库级别的配置文件**, 打开看看

![[Pasted image 20250731104449.png]]
可以看到仓库级别的配置文件中并**没有** user.name 或者 user.email

让我们看看 `~/.gitconfig` 文件, 这是 linux **用户级别的配置文件**

![[Pasted image 20250731104639.png]]

这说明 git 配置是会继承的

### `.git/objects` 文件夹

这个文件夹下存储了所有的 git 存储对象

### `.git/refs` 文件夹

这个文件夹下存储了当前 `HEAD` 指针相关的文件