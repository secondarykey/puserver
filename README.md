# UML

PlantUML を利用してUMLを編集するHTTPサーバです。
サーブレットコンテナを準備するのが大変だったので作成してみました。

# Install

まだ固定設定していますが、手元にPlantUMLのjarを配置して実行
javaがあることが前提です。

# あとがき

作っておいて、結局javaのインストールに依存してしまうので
コンテナ準備して、plantuml-serverを利用したほうがいいかな、、、とも思ったりした。
JNIとか利用してjava自体を排除しないといけない気がしているけど、
技術的にできるのか以上にライセンス問題ないかも確認

# Issue

- 生成部分をプログレスで見えるようにする

- 本格的に作るならReact化
- ユーザ管理（グループ、ロール設計）
- ドキュメント管理(ユーザに紐づいた文書一覧 少なくとも下書きと清書を作成)
- JNIによるjavaコマンドの排除
- WYSIWYGエディター化(そもそもあまり知らないので入力して気づいたら)
- Excel 化？(Excelのセルに書いてURLで画像生成)
