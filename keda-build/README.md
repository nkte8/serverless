# 本イメージについて  

以下のプロジェクトをdockerコンテナ内でコンパイルするdockerfileです。  
https://github.com/kedacore/keda

## ビルド方法  
build方法およびdockerfileを参考に作成します。
```sh
make build ### 全てのコンポーネントをビルド
make manager ### keda coreのビルド
make adapter ### keda metrics serverのビルド
```
https://github.com/kedacore/keda/blob/main/BUILD.md  
https://github.com/kedacore/keda/blob/main/Dockerfile  
https://github.com/kedacore/keda/blob/main/Dockerfile.adapter  