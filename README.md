# nannan

# 缓存设置
## 文件配置
> 
## memcache 配置
> os.Setenv("key","memcache://")
## redis 配置
> os.Setenv("key","redis://user:passwd@ip:port/?retries=3&db=1")
## redis 群集配置
> os.Setenv("key","redis_cluster://"))
## redis 哨兵配置
> os.Setenv("key","redis_sentinel://user:passwd@ip:port/?retries=3&db=1&servers=ip:port,ip:port"))