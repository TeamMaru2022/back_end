# back-end

## センサーとバックエンドのフロー

!(blob:https://teams.microsoft.com/3fa80bf1-b6d1-4b2d-9f05-78a2b5b04cc0)



## 起動方法

※windowsは、docker for windows を起動してから行う。

**[windowsの場合]**

**back-end/db/my.conf をエクスプローラーのプロパティから読み取り専用に変更しておくこと。**

**→JsonファイルやDBから取得した値が文字化けしてしまうため。**

1. dockerを起動させる

```
docker-compose up
```

初回はエラーが発生して起動できない可能性あり。

その際は、

2. dockerを停止する

```
docker-compose down
```

をして、再度起動させる。
