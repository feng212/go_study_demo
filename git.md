---
typora-copy-images-to: images
---

学习使用git对于职场人来说是一项基本的技能，在此我只展示在工作过程中经常使用到的命令。

推荐课程：

菜鸟教程：https://www.runoob.com/git/git-tutorial.html

git 官 网 ：https://git-scm.com/book/en/v2

在git官网学习过之后你会有基本的概念和使用基础，通过菜鸟教程提供的详细案例，可以快速使用。



### 1.clone  克隆

直接clone项目: 默认主分支(master)

```git
git clone  origin
git clone  https://gitee.com/ojojo5555555/go_study_demo.git
```

指定分支clone: 使用-b命令选择分支

```
git clone -b branch origin
git clone -b 220503_init_data_type  https://gitee.com/ojojo5555555/go_study_demo.git
```

### 2.remote 远程仓库

添加远程仓库 : add命令 ，可给远程仓库起个名称

```
git remote add name origin
git remote add mine https://gitee.com/ojojo5555555/go_study_demo.git
```

修改远程仓库名称

```
git remote add old_name new_name
git remote mine mine Gitee
```

删除远程仓库

```
git remote rm name
```

显示所有远程仓库

```
git remote -v
```

### 3.branch 分支：

每个分支都有不同的内容，从不同的分支进行切换都会复制出一份与之相同的本地分支

创建本地分支

```
git branch name
```

切换分支

```
git checkout checkout_name(切换的分支名称)
```

删除分支：-d命令

```
git branch -d name
```

当clone下来后想快速创建并切换远程分支，可用如下命令：

-b 指定远程分支时使用，本地分支不需要此参数

```
git checkout -b new_branch_name checkout_name
				(切换后的新本支)   (要切换的远程分支)
```

显示分支:

-a 所有分支,包括远程分支

默认本地分支

```
git branch -a
```



remote和branch信息，可在项目目录下的.git文件的config文件中查看

如：

[core]
	repositoryformatversion = 0
	filemode = false
	bare = false
	logallrefupdates = true
	ignorecase = true
[submodule]
	active = .
[remote "origin"]
	url = https://gitee.com/ojojo5555555/geias.git
	fetch = +refs/heads/*:refs/remotes/origin/*
[branch "master"]
	remote = origin
	merge = refs/heads/master



### 4.status 当前仓库状态

```
git status 显示修改后的所有文件
```

### 5.add 添加到暂存区

```
git add .  所有都到暂存区
git add name 指定文件到暂存区，name包括文件类型
```

### 6.commit 提交暂存到本地仓库

```
git commit -m 描述
```

### 7.push 提交到远程仓库

```
git push remote_name branch_name
		(远程仓库名称) （此时的本地分支名称，若远程不存在该分支，push之后远程会自动创建该分支） 
		remote add起时的名称
```

### 8.pull合并代码

**git pull** 其实就是 **git fetch** 和 **git merge FETCH_HEAD** 的简写。 命令格式如下：

```
git pull <远程主机名> <远程分支名>:<本地分支名>
```

