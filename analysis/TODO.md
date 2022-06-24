# Setup
- Wrap backend with nginx to utilize alp
- Make `test.sh` for a bench and analysis iteration
- Use golangci-lint

# MEMO
- 113 isu
- 117k isu_condition
- 25 distinct characters
- isu_condition timestamp format: '2021-08-10 20:07:03'

# Score
- 1. Default 3000
- 2. 1 + Index in isu_condition + No select image in /api/isu : 16000
- 3. 2 + Cache image (/icon)

# Idea
- (done, highly effective) Use Nginx cache for asset and/or set Cache-Control header
- Make conditionLevel separate row. Use WHERE IN clause for /api/condition/xxx
- (done, highly effective) Use WHERE for timestamp in /api/isu/graph
- (done, highly effective) Bulk insert condition hourly
- Bulk insert in larger chunk to deal with more users
- (tested, This results in huge number of users) Use Nginx cache for /api/trend
- (tested, same as above) Cache latest condition for GET /api/condition
- Use limit 20 for each condition level in /api/condition
- (Done) Profile backend
- Profile db
- Use db connection pool
- (Done) No use row.next() in getIsuGraph

# Bug note
- Setting `POST_ISUCONDITION_TARGET_BASE_URL` to port 3001 resulted in 0 score. I must set `ISUXBENCH_ALL_ADDRESSES` to nginx, too.

# Must do before end
- Disable db log
- Disable golang log
- Disable nginx log
- Check https://blog.recruit.co.jp/rtc/2021/04/26/isucon-2021-winter/