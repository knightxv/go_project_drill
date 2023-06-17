
if [ "$(uname)" == "Darwin" ]; then
 PROJECTROOT=/Users/apple/Documents/Open-IM-Server
 WorkProject=/Users/apple/Documents/gosrc/src/Open-IM-Server
 /bin/cp -rf ../pkg/base_info $PROJECTROOT/pkg/
 /bin/cp -rf ../pkg/common/db $PROJECTROOT/pkg/common/
 /bin/cp -r ../internal/api   $PROJECTROOT/internal/
 /bin/cp -r ../cmd/user_score_api/main.go   $PROJECTROOT/cmd/user_score_api/main.go
 cd $PROJECTROOT
fi

if [[ "$(uname)" == "MINGW"* ]]; then
  SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
  echo $SCRIPT_DIR
  cd SCRIPT_DIR/../ # 进入工作目录
fi

swag init --parseVendor --parseInternal --parseDependency --parseDepth 100   -o ./cmd/user_score_api/docs  -g ./cmd/user_score_api/main.go
if [ "$(uname)" == "Darwin" ]; then
 /bin/cp -rf  $PROJECTROOT/cmd/user_score_api/docs   $WorkProject/cmd/user_score_api/
 cd $WorkProject/script
fi
