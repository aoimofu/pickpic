# pickpicコマンド
Twitter IDを引数に渡すと、最新の投稿から画像のURLをスクレイピングするコマンドです。

## ビルド ( CentOS7 )
```
# goqueryのインストール
go get github.com/PuerkitoBio/goquery

# 本体をclone
git clone git@github.com:aoimofu/pickpic.git

# ビルド実行
cd pickpic/
go build
```

## 使い方
```
pickpic -i <TwitterID>
```

## 例
```
pickpic -i anju_inami

## 以下結果
https://pbs.twimg.com/media/DVwpKmKU0AARmEE.jpg
https://pbs.twimg.com/media/DVjx6XcVAAAQho9.jpg
https://pbs.twimg.com/media/DVjx6XbV4AEvKo4.jpg
https://pbs.twimg.com/media/DVbmt_FVAAEJBcY.jpg
https://pbs.twimg.com/media/DVbFYs6U0AAbLnn.jpg
(以下省略)
```
