		bzcat                        # 查看bzip2
		gzip A                       # 直接压缩文件 # 压缩后源文件消失
		gunzip A.gz                  # 直接解压文件 # 解压后源文件消失
		gzip -dv 1.tar.gz            # 解压gzip到tar
		gzip -v 1.tar                # 压缩tar到gz
		unzip zip.zip                # 解压zip
		zip zip.zip *                # 压缩zip
		# rar3.6下载:  http://www.rarsoft.com/rar/rarlinux-3.6.0.tar.gz
		rar a rar.rar *.jpg          # 压缩文件为rar包
		unrar x rar.rar              # 解压rar包
		7z a 7z.7z *                 # 7z压缩
		7z e 7z.7z                   # 7z解压

	}
	
	文件ACL权限控制{

		getfacl 1.test                      # 查看文件ACL权限
		setfacl -R -m u:xuesong:rw- 1.test  # 对文件增加用户的读写权限 -R 递归

	}
	
	svn更新代码{

		--force # 强制覆盖
		/usr/bin/svn --username user --password passwd co  $Code  ${SvnPath}src/                 # 检出整个项目
		/usr/bin/svn --username user --password passwd export  $Code$File ${SvnPath}src/$File    # 导出个别文件

	}

	恢复rm删除的文件{

		# debugfs针对 ext2   # ext3grep针对 ext3   # extundelete针对 ext4
		df -T   # 首先查看磁盘分区格式
		umount /data/     # 卸载挂载,数据丢失请首先卸载挂载,或重新挂载只读
		ext3grep /dev/sdb1 --ls --inode 2         # 记录信息继续查找目录下文件inode信息
		ext3grep /dev/sdb1 --ls --inode 131081    # 此处是inode
		ext3grep /dev/sdb1 --restore-inode 49153  # 记录下inode信息开始恢复目录

	}
	
}

2软件{

	rpm{

		rpm -ivh lynx          # rpm安装
		rpm -e lynx            # 卸载包
		rpm -e lynx --nodeps   # 强制卸载
		rpm -qa                # 查看所有安装的rpm包
		rpm -qa | grep lynx    # 查找包是否安装

	最大连接数{

		ulimit -SHn 65535  # 修改最大打开文件数(等同最大连接数)
		ulimit -a          # 查看
		
		/etc/security/limits.conf         # 进程最大打开文件数
		# nofile 可以被理解为是文件句柄数 文件描述符 还有socket数
		* soft nofile 65535
		* hard nofile 65535
		# 最大进程数
		* soft nproc 65535
		* hard nproc 65535

		# 如果/etc/security/limits.d/有配置文件，将会覆盖/etc/security/limits.conf里的配置
		# 即/etc/security/limits.d/的配置文件里就不要有同样的参量设置
		/etc/security/limits.d/90-nproc.conf    # centos6.3的最大进程数文件
		* soft nproc 65535       
		* hard nproc 65535

	}
	
	sudo{

		visudo     # sudo命令权限添加
		用户  别名(可用all)=NOPASSWD:命令1，命令2
		wangming linuxfan=NOPASSWD:/sbin/apache start,/sbin/apache restart
		UserName ALL=(ALL) ALL
		peterli        ALL=(ALL)       NOPASSWD:/sbin/service
		Defaults requiretty             # sudo不允许后台运行,注释此行既允许
		Defaults !visiblepw             # sudo不允许远程,去掉!既允许

	}

	grub开机启动项添加{

		vim /etc/grub.conf
		title ms-dos
		rootnoverify (hd0,0)
		chainloader +1

	}

	stty{

		#stty时一个用来改变并打印终端行设置的常用命令

		stty iuclc          # 在命令行下禁止输出大写
		stty -iuclc         # 恢复输出大写
		stty olcuc          # 在命令行下禁止输出小写
		stty -olcuc         # 恢复输出小写
		stty size           # 打印出终端的行数和列数
		stty eof "string"   # 改变系统默认ctrl+D来表示文件的结束 
		stty -echo          # 禁止回显

			# 格式： mysql -h 主机地址 -u 用户名 -p 用户密码
			mysql -h110.110.110.110 -P3306 -uroot -p
			mysql -uroot -p -S /data1/mysql5/data/mysql.sock -A  --default-character-set=GBK

		}
		
		shell执行mysql命令{

			mysql -u$username -p$passwd -h$dbhost -P$dbport -A -e "      
			use $dbname;
			delete from data where date=('$date1');
			"    # 执行多条mysql命令
			mysql -uroot -p -S mysql.sock -e "use db;alter table gift add column accountid  int(11) NOT NULL;flush privileges;"    # 不登陆mysql插入字段

		}

		备份数据库{

			mysqldump -h host -u root -p --default-character-set=utf8 dbname >dbname_backup.sql               # 不包括库名，还原需先创建库，在use 
			mysqldump -h host -u root -p --database --default-character-set=utf8 dbname >dbname_backup.sql    # 包括库名，还原不需要创建库
			/bin/mysqlhotcopy -u root -p    # mysqlhotcopy只能备份MyISAM引擎
			mysqldump -u root -p -S mysql.sock --default-character-set=utf8 dbname table1 table2  > /data/db.sql    # 备份表
			mysqldump -uroot -p123  -d database > database.sql    # 备份数据库结构
			
			innobackupex --user=root --password="" --defaults-file=/data/mysql5/data/my_3306.cnf --socket=/data/mysql5/data/mysql.sock --slave-info --stream=tar --tmpdir=/data/dbbackup/temp /data/dbbackup/ 2>/data/dbbackup/dbbackup.log | gzip 1>/data/dbbackup/db50.tar.gz   # xtrabackup备份需单独安装软件 优点: 速度快,压力小,可直接恢复主从复制

		}
		 
		还原数据库{

			mysql -h host -u root -p dbname < dbname_backup.sql   
			source 路径.sql   # 登陆mysql后还原sql文件

		}
		
		赋权限{

			# 指定IP: $IP  本机: localhost   所有IP地址: %   # 通常指定多条
			grant all on zabbix.* to user@"$IP";             # 对现有账号赋予权限
			grant select on database.* to user@"%" Identified by "passwd";     # 赋予查询权限(没有用户，直接创建)
			grant all privileges on database.* to user@"$IP" identified by 'passwd';         # 赋予指定IP指定用户所有权限(不允许对当前库给其他用户赋权限)
			grant all privileges on database.* to user@"localhost" identified by 'passwd' with grant option;   # 赋予本机指定用户所有权限(允许对当前库给其他用户赋权限)
			grant select, insert, update, delete on database.* to user@'ip'identified by "passwd";   # 开放管理操作指令
			revoke all on *.* from user@localhost;     # 回收权限

		}

		更改密码{

			update user set password=password('apzmnr') where user='root'
			mysqladmin -u root password 'xuesong'

		}

		mysql忘记密码后重置{

			cd /data/mysql5
			/data/mysql5/bin/mysqld_safe --user=mysql --skip-grant-tables --skip-networking &
			update user set password=password('944593') where user='root';

		}
		
		mysql主从复制失败恢复{

			slave stop;
			reset slave;
			change master to master_host='10.10.10.110',master_port=3306,master_user='repl',master_password='oygl',master_log_file='master-bin.000010',master_log_pos=107,master_connect_retry=60;
			slave start;

		}
		
		检测mysql主从复制延迟{
			
			1、在从库定时执行更新主库中的一个timeout数值
			2、同时取出从库中的timeout值对比判断从库与主库的延迟
		
		}
	}

	mongodb{

		一、启动{
		
			# 不启动认证
			./mongod --port 27017 --fork --logpath=/opt/mongodb/mongodb.log --logappend --dbpath=/opt/mongodb/data/
			# 启动认证
			./mongod --port 27017 --fork --logpath=/opt/mongodb/mongodb.log --logappend --dbpath=/opt/mongodb/data/ --auth

			# 配置文件方式启动
			cat /opt/mongodb/mongodb.conf
			  port=27017                       # 端口号
			  fork=true                        # 以守护进程的方式运行，创建服务器进程
		export JAVA_HOME=/usr/local/jdk1.5.0_14 
		export CLASSPATH=.:$JAVA_HOME/jre/lib/rt.jar:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar 
		export PATH=$PATH:$JAVA_HOME/bin
		. /etc/profile

	}

}

5网络{

	rz   # 通过ssh上传小文件
	sz   # 通过ssh下载小文件
	ifconfig eth0 down                  # 禁用网卡
	ifconfig eth0 up                    # 启用网卡
	ifup eth0:0                         # 启用网卡
	mii-tool em1                        # 查看网线是否连接
	traceroute www.baidu.com            # 测试跳数
	vi /etc/resolv.conf                 # 设置DNS  nameserver IP 定义DNS服务器的IP地址
	nslookup www.moon.com               # 解析域名IP
	dig -x www.baidu.com                # 解析域名IP
	curl -I www.baidu.com               # 查看网页http头
	tcpdump tcp port 22                 # 抓包
	lynx                                # 文本上网
	wget -P 路径 http地址               # 下载  包名:wgetrc
	curl -d "user=xuesong&pwd=123" http://www.abc.cn/Result    # 提交web页面表单 需查看表单提交地址
	rsync -avzP -e "ssh -p 22" /dir user@$IP:/dir              # 同步目录 # --delete 无差同步 删除目录下其它文件
	ifconfig eth0:0 192.168.1.221 netmask 255.255.255.0        # 增加逻辑IP地址
	mtr -r www.baidu.com                                       # 测试网络链路节点响应时间 # trace ping 结合
	echo 1 > /proc/sys/net/ipv4/icmp_echo_ignore_all           # 禁ping
	ipcalc -m "$ip" -p "$num"                                  # 根据IP和主机最大数计算掩码
	dig +short txt hacker.wp.dg.cx                             # 通过 DNS 来读取 Wikipedia 的hacker词条
	host -t txt hacker.wp.dg.cx                                # 通过 DNS 来读取 Wikipedia 的hacker词条
	net rpc shutdown -I IP_ADDRESS -U username%password        # 远程关掉一台WINDOWS机器
	wget --random-wait -r -p -e robots=off -U Mozilla www.example.com    # 递归方式下载整个网站
	
	netstat{

		-a     # 显示所有连接中的Socket
		-t     # 显示TCP连接
		-u     # 显示UDP连接
		-n     # 显示所有已建立的有效连接
		netstat -anlp           # 查看链接
		netstat –r              # 查看路由表

	}
	
	ssh{

		ssh -p 22 user@192.168.1.209                            # 从linux ssh登录另一台linux 
		ssh -p 22 root@192.168.1.209 CMD                        # 利用ssh操作远程主机
		scp -P 22 文件 root@ip:/目录                            # 把本地文件拷贝到远程主机
		sshpass -p '密码' ssh -n root@$IP "echo hello"          # 指定密码远程操作
		ssh -o StrictHostKeyChecking=no $IP                     # ssh连接不提示yes
		ssh -t "su -"                                           # 指定伪终端 客户端以交互模式工作
		scp root@192.168.1.209:远程目录 本地目录                # 把远程指定文件拷贝到本地
		ssh -N -L2001:remotehost:80 user@somemachine            # 用SSH创建端口转发通道
		ssh -t host_A ssh host_B                                # 嵌套使用SSH
		ssh -t -p 22 $user@$Ip /bin/su - root -c {$Cmd};        # 远程su执行命令 Cmd="\"/sbin/ifconfig eth0\""
		ssh-keygen -t rsa                                       # 生成密钥
		ssh-copy-id -i xuesong@10.10.10.133                     # 传送key
		vi $HOME/.ssh/authorized_keys                           # 公钥存放位置
		sshfs name@server:/path/to/folder /path/to/mount/point  # 通过ssh挂载远程主机上的文件夹
		fusermount -u /path/to/mount/point                      # 卸载ssh挂载的目录
		ssh user@host cat /path/to/remotefile | diff /path/to/localfile -                # 用DIFF对比远程文件跟本地文件
		su - user -c "ssh user@192.168.1.1 \"echo -e aa |mail -s test mail@163.com\""    # 切换用户登录远程发送邮件

	}

	网卡配置文件{

		vi /etc/sysconfig/network-scripts/ifcfg-eth0

		DEVICE=eth0
		BOOTPROTO=none
		BROADCAST=192.168.1.255
		HWADDR=00:0C:29:3F:E1:EA
		IPADDR=192.168.1.55
		sed -i -e '1,3d' -e 's/1/2/'                  # 多重编辑(先删除1-3行，在将1替换成2)
		sed -e 's/@.*//g' -e '/^$/d'                  # 删除掉@后面所有字符，和空行
		sed -n -e "{s/文本(正则)/替换的内容/p}"       # 替换并打印出替换行
		sed -n -e "{s/^ *[0-9]*//p}"                  # 打印并删除正则表达式的那部分内容
		echo abcd|sed 'y/bd/BE/'                      # 匹配字符替换
		sed '/^#/b;y/y/P/' 2                          # 非#号开头的行替换字符
		sed '/suan/r 读入文件'                        # 找到含suan的行，在后面加上读入的文件内容
		sed -n '/no/w 写入文件'                       # 找到含no的行，写入到指定文件中
		sed '/regex/G'                                # 在匹配式样行之后插入一空行
		sed '/regex/{x;p;x;G;}'                       # 在匹配式样行之前和之后各插入一空行
		sed 'n;d'                                     # 删除所有偶数行
		sed 'G;G'                                     # 在每一行后面增加两空行
		sed '/^$/d;G'                                 # 在输出的文本中每一行后面将有且只有一空行
		sed 'n;n;n;n;G;'                              # 在每5行后增加一空白行
		sed -n '5~5p'                                 # 只打印行号为5的倍数
		seq 1 30|sed  '5~5s/.*/a/'                    # 倍数行执行替换
		sed -n '3,${p;n;n;n;n;n;n;}'                  # 从第3行开始，每7行显示一次
		sed -n 'h;n;G;p'                              # 奇偶调换
		seq 1 10|sed '1!G;h;$!d'                      # 倒叙排列
		ls -l|sed -n '/^.rwx.*/p'                     # 查找属主权限为7的文件
		sed = filename | sed 'N;s/\n/\t/'             # 为文件中的每一行进行编号(简单的左对齐方式)
		sed 's/^[ \t]*//'                             # 将每一行前导的"空白字符"(空格，制表符)删除,使之左对齐 
		sed 's/^[ \t]*//;s/[ \t]*$//'                 # 将每一行中的前导和拖尾的空白字符删除
		echo abcd\\nabcde |sed 's/\\n/@/g' |tr '@' '\n'        # 将换行符转换为换行
		cat tmp|awk '{print $1}'|sort -n|sed -n '$p'           # 取一列最大值
		sed -n '{s/^[^\/]*//;s/\:.*//;p}' /etc/passwd          # 取用户家目录(匹配不为/的字符和匹配:到结尾的字符全部删除)
		sed = filename | sed 'N;s/^/      /; s/ *\(.\{6,\}\)\n/\1   /'   # 对文件中的所有行编号(行号在左，文字右端对齐)
		/sbin/ifconfig |sed 's/.*inet addr:\(.*\) Bca.*/\1/g' |sed -n '/eth/{n;p}'   # 取所有IP

		修改keepalive配置剔除后端服务器{

			sed -i '/real_server.*10.0.1.158.*8888/,+8 s/^/#/' keepalived.conf
			sed -i '/real_server.*10.0.1.158.*8888/,+8 s/^#//' keepalived.conf

		}
		
		模仿rev功能{

			echo 123 |sed '/\n/!G;s/\(.\)\(.*\n\)/&\2\1/;//D;s/.//;'
			/\n/!G;         　　　　　　# 没有\n换行符，要执行G,因为保留空间中为空，所以在模式空间追加一空行
			s/\(.\)\(.*\n\)/&\2\1/;     # 标签替换 &\n23\n1$ (关键在于& ,可以让后面//匹配到空行)
			//D;            　　　　　　# D 命令会引起循环删除模式空间中的第一部分，如果删除后，模式空间中还有剩余行，则返回 D 之前的命令，重新执行，如果 D 后，模式空间中没有任何内容，则将退出。  //D 匹配空行执行D,如果上句s没有匹配到,//也无法匹配到空行, "//D;"命令结束
			s/.//;          　　　　　　# D结束后,删除开头的 \n

		}

	}

	dialog菜单{
	
		# 默认将所有输出用 stderr 输出，不显示到屏幕   使用参数  --stdout 可将选择赋给变量
				switch $1 {
					Ssh_Cmd {
						spawn ssh -t -p $Port -o StrictHostKeyChecking=no $RemoteUser@$Ip /bin/su - root -c \\\"$Cmd\\\"
						jiaohu
					}
					Ssh_Script {
						spawn scp -P $Port -o StrictHostKeyChecking=no $ScriptPath $RemoteUser@$Ip:/tmp/${ScriptPath##*/};
						jiaohu
						spawn ssh -t -p $Port -o StrictHostKeyChecking=no $RemoteUser@$Ip /bin/su - root -c  \\\"/bin/sh /tmp/${ScriptPath##*/}\\\" ;
						jiaohu
					}
					Scp_File {
						spawn scp -P $Port -o StrictHostKeyChecking=no -r $ScpPath $RemoteUser@$Ip:${ScpRemotePath};
						jiaohu
					}
				}
				"

			}

			交互双引号引用较长变量{
			
				#!/bin/bash
				RemoteUser=xuesong12
				Ip=192.168.1.2
				RemotePasswd=fyuovni
				Cmd="/bin/echo "$PubKey" > "$RemoteKey"/authorized_keys"

				/usr/bin/expect -c "
				set timeout 10
				spawn ssh -o StrictHostKeyChecking=no $RemoteUser@$Ip {$Cmd};
				expect {
					password: {
						send_user RemotePasswd\n
						send ${RemotePasswd}\r;
						interact;
					}
					eof {
						send_user eof\n
					}
				}
				"

			}

			telnet{
			
				#!/bin/bash
				Ip="10.0.1.53"
				a="\{\'method\'\:\'doLogin\'\,\'params\'\:\{\'uName\'\:\'bobbietest\'\}"
				/usr/bin/expect -c"

	}

	找出出现次数最多{

		awk '{print $1}' file|sort |uniq -c|sort -k1r

	}
	
	判断脚本参数是否正确{

		./test.sh  -p 123 -P 3306 -h 127.0.0.1 -u root
		#!/bin/sh
		if [ $# -ne 8 ];then
			echo "USAGE: $0 -u user -p passwd -P port -h host"
			exit 1
		fi

		while getopts :u:p:P:h: name
		do
			case $name in
			u)
				mysql_user=$OPTARG
			;;
			p)
				mysql_passwd=$OPTARG
			;;
			P)
				mysql_port=$OPTARG
			;;
			h)
				mysql_host=$OPTARG
			;;
			*)
				echo "USAGE: $0 -u user -p passwd -P port -h host"
				exit 1
			;;
			esac
		done

		if [ -z $mysql_user ] || [ -z $mysql_passwd ] || [ -z $mysql_port ] || [ -z $mysql_host ]
		then
			echo "USAGE: $0 -u user -p passwd -P port -h host"
			exit 1
		fi

		echo $mysql_user $mysql_passwd $mysql_port  $mysql_host
		#结果 root 123 3306 127.0.0.1
	
	}


	两日期间隔天数{
	
		D1=`date -d '20070409' +"%s"`
		D2=`date -d '20070304 ' +"%s"`
		D3=$(($D1 - $D2))
		echo $(($D3/60/60/24)) 
		
	}

	while执行ssh只循环一次{
		
		cat -    # 让cat读连接文件stdin的信息
		seq 10 | while read line; do ssh localhost "cat -"; done        # 显示的9次是被ssh吃掉的
		seq 10 | while read line; do ssh -n localhost "cat -"; done     # ssh加上-n参数可避免只循环一次
		
	}

	ssh批量执行命令{

		#版本1
		#!/bin/bash
		while read line
		do 
		Ip=`echo $line|awk '{print $1}'`
		Passwd=`echo $line|awk '{print $2}'`
		ssh -n localhost "cat -"
		sshpass -p "$Passwd" ssh -n -t -o StrictHostKeyChecking=no root@$Ip "id"  
		done<iplist.txt

		#版本2
		#!/bin/bash
		Iplist=`awk '{print $1}' iplist.txt`
		for Ip in $Iplist
		do
		Passwd=`awk '/'$Ip'/{print $2}' iplist.txt`
		sshpass -p "$Passwd" ssh -n -t -o StrictHostKeyChecking=no root@$Ip "id" 
		done

	}

	在同一位置打印字符{

		#!/bin/bash
		echo -ne "\t"
		for i in `seq -w 100 -1 1`
		do
			echo -ne "$i\b\b\b";      # 关键\b退格
			sleep 1;
		done

	}

	多进程后台并发控制{

		#!/bin/bash
		test () {
			echo $a
			sleep 5
		}
		for a in `seq 1 30`
		do