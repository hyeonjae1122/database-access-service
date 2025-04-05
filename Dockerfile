# 1. Go 언어 베이스 이미지 사용
FROM golang:1.24

# 2. 작업 디렉토리 설정
WORKDIR /app

# 3. Go 소스코드 복사
COPY go.mod .
COPY main.go .

# 4. 빌드
RUN go build -o server

# 5. 컨테이너에서 실행될 기본 명령어 설정
CMD ["./server"]

# 6. 포트 노출
EXPOSE 8080