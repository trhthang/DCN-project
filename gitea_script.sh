#!/bin/bash

# TOKEN="0f20cbb332d816d8ae345987bda130d00a7f711e"
# USERNAME="trhthang"
# REPO_NAME="edge03"

# curl -X POST -H "Authorization: token $TOKEN" -H "Content-Type: application/json" -d "{\"name\":\"$REPO_NAME\"}" http://192.168.11.92:30357/api/v1/user/repos

# -------------------------------------------------xóa repo:

#!/bin/bash

# TOKEN="0f20cbb332d816d8ae345987bda130d00a7f711e"
# USERNAME="trhthang"
# REPO_NAME="$1"

# # Xóa repository
# curl -X DELETE -H "Authorization: token $TOKEN" http://192.168.11.92:30357/api/v1/repos/$USERNAME/$REPO_NAME



# tạo repo với script:

# #!/bin/bash

# TOKEN="0f20cbb332d816d8ae345987bda130d00a7f711e"
# USERNAME="trhthang"
# REPO_NAME="$1"

# curl -X POST -H "Authorization: token $TOKEN" -H "Content-Type: application/json" -d "{\"name\":\"$REPO_NAME\"}" http://192.168.11.92:30357/api/v1/user/repos

# --------------------------------------------tạo branch

!/bin/bash

TOKEN="0f20cbb332d816d8ae345987bda130d00a7f711e"
USERNAME="trhthang"
REPO_NAME="$1"

# Kiểm tra xem có tham số được truyền vào không
if [ -z "$REPO_NAME" ]; then
  echo "Usage: $0 <repository_name>"
  exit 1
fi

# Tạo repository
curl -X POST -H "Authorization: token $TOKEN" -H "Content-Type: application/json" -d "{\"name\":\"$REPO_NAME\"}" http://192.168.11.92:30357/api/v1/user/repos

# Tạo nội dung README
README_CONTENT="This is the README file for the repository $REPO_NAME."

# Tạo tệp README
curl -X POST -H "Authorization: token $TOKEN" -H "Content-Type: application/json" -d "{\"content\":\"$(echo -n $README_CONTENT | base64)\"}" http://192.168.11.92:30357/api/v1/repos/$USERNAME/$REPO_NAME/contents/README.md

# Tạo branch "main" (hoặc master) - đối với Gitea, có thể sử dụng "master" thay vì "main"
curl -X POST -H "Authorization: token $TOKEN" -H "Content-Type: application/json" -d "{\"ref_name\":\"main\"}" http://192.168.11.92:30357/api/v1/repos/$USERNAME/$REPO_NAME/git/refs

# Lấy SHA của commit cuối cùng trên branch "main" (hoặc master)
# SHA=$(curl -s -H "Authorization: token $TOKEN" http://192.168.11.92:30357/api/v1/repos/$USERNAME/$REPO_NAME/commits/main | jq -r '.sha')
SHA=$(curl -s -H "Authorization: token $TOKEN" http://192.168.11.92:30357/api/v1/repos/$USERNAME/$REPO_NAME/commits/main)

# Đẩy README lên branch "main" (hoặc master)
curl -X POST -H "Authorization: token $TOKEN" -H "Content-Type: application/json" -d "{\"ref\":\"refs/heads/main\",\"sha\":\"$SHA\"}" http://192.168.11.92:30357/api/v1/repos/$USERNAME/$REPO_NAME/git/refs



# Đăng ký với Porch
kpt alpha repo register \
  --namespace default \
  --repo-basic-username=$USERNAME \
  --repo-basic-password=$TOKEN \
  --deployment=true \
  http://192.168.11.92:30357/trhthang/$REPO_NAME.git