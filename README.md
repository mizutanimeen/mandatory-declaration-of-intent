# mandatory-declaration-of-intent

## 概要

意見出しの際に人の意見を見て悪い方向に自分の意見が変わってしまったり

自分から意見を出さずに人の意見を見てばかりになることを防止する web アプリ

意見を出さないと他の人の意見が見れない

### VERSION 0.0.0

## 開発手順

- 権限とか特になし
- ルームごとにパスワードを付けれる
- ユーザーを作成でき、ユーザーごとに許可を出せる

### メモ

https://app.diagrams.net/#G1wSNS1Fk22s8tYFcnFMgeRdeICfvrPR0h

react 側で get room すると cookie 名とかも手に入っちゃうから ID、名前、説明のみ取得できる API にする。

HttpOnly: true,を付けても動作するか

