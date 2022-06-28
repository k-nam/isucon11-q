# Setup
- Wrap backend with nginx to utilize alp
- Make `test.sh` for a bench and analysis iteration
- Use golangci-lint

# MEMO
- 113 isu
- 117k isu_condition
- 25 distinct characters
- isu_condition timestamp format: '2021-08-10 20:07:03'

# Idea (easy)
- (done, highly effective) Use Nginx cache for asset and icon, and/or set Cache-Control header
- (Done, highly effective) Select only necessary cols
- (Done, highly effective) Use WHERE for timestamp in /api/isu/graph
- (Done, highly effective) Use WHERE IN for condition in /api/condition
- (Done, highly effective) Use upstream keepalive in Nginx
- (Done, highly effective) Use db connection pool

# Using memory
- (Done, highly effective) Bulk insert condition hourly
- (Done, highly effective) Use cached latestCondition in getIsuList
- (Done, highly effective) Cache users in getUserFromSession

# Can't use
- (Done, cannot pass benchmark when I change table schema) Make conditionLevel separate row. Use WHERE IN clause for /api/condition/xxx
- (tested, This results in huge number of users) Use Nginx cache for /api/trend
- (tested, same as above) Cache latest condition for GET /api/condition

# Chores
- (Done) Profile backend
- (Done) Profile db
- (Done, no effect) No use row.next() in getIsuGraph
- (Done, NG) Try dealing only with info level conditions

# TODO
- ( ) Make isu_condition table's index more effective by compress condition column

# Bug note
- Setting `POST_ISUCONDITION_TARGET_BASE_URL` to port 3001 resulted in 0 score. I must set `ISUXBENCH_ALL_ADDRESSES` to nginx, too.

# Must do before end
- Disable db log
- Disable golang log
- Disable nginx log
- Check https://blog.recruit.co.jp/rtc/2021/04/26/isucon-2021-winter/