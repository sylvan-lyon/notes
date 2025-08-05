同样, 重新创建一个仓库, 并提交一个文件. 同样, 为了少看两个文件, 我们删除 `.sample` 文件

![[Pasted image 20250802141406.png]]

完成后仓库目录大致如下:

![[Pasted image 20250802141437.png]]

接下来, 就是今天的重头戏啦

# git branch

## 执行命令

### 执行 `git branch` 查看分支

![[Pasted image 20250802141541.png]]
### 在当前分支上创建一个新的分支, 并切换分支

![[Pasted image 20250802141802.png]]

> [!info]
> 使用 `git branch <分支名>` 创建一个分支
> 
> 使用 `git checkout <分支名>` 切换到指定分支
> 
> `git branch` 中, 高亮且有星号的分支是当前分支, 你也能在我的提示符中看到

## 查看变化

![[Pasted image 20250802142230.png]]

这就是在创建分支之后的文件夹结构, 需要注意的是, objects 文件夹并没有什么改动, 但是 `refs/heads` 目录下多了一个 `dev` 文件, 我们来看看

![[Pasted image 20250802142439.png]]
![[Pasted image 20250802142535.png]]

我们可以看到这个目录下的两个文件的内容是一样的, 都指向同一个 `commit` 对象

在上一节中, 我们讨论一下 `HEAD` 这个文件, 让我们来看一下

![[Pasted image 20250802142728.png]]

注意到 `.git/HEAD` 文件夹指向当前分支的 `refs/heads/<分支名>` 的文件

于是稍微思考一下, 就能想出切换分支的时候工作区为什么能够正确的将分支中的进度取出来的了.

1. `git checkout <分支名>` 将 `HEAD` 指针指向对应分支的 `head` 文件
2. `git checkout <分支名>` 分析目标分支中的 `commit` 对象, 获取该 `commit` 对象对应的 `tree` 对象树, 通过 `tree` 对象获取 `blob` 对象, 将 `blob` 对象解压到对应的位置, 分支切换完成