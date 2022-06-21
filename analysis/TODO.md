# Setup
- Wrap backend with nginx to utilize alp
- Make `test.sh` for a bench and analysis iteration
- Use golangci-lint

# MEMO
- 113 isu
- 117k isu_condition
- 25 distinct characters


# Score
- 1. Default 3000
- 2. 1 + Index in isu_condition + No select image in /api/isu : 16000
- 3. 2 + Cache image (/icon)



# Bug note
- Setting `POST_ISUCONDITION_TARGET_BASE_URL` to port 3001 resulted in 0 score