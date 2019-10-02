# Portfolio

## はじめに
こんにちは。閲覧いただきありがとうございます。

現在転職活動用にポートフォリオを作成しております。
「仮想通貨疑似売買アプリ」を作成しようとしています。

このポートフォリオは現在作成中であり、未完成となっています。

今までの勉強記録は**やってきたこと、苦戦したこと**に記載しております。  随時更新中です。

## システム構成
- 使用言語 
Vue.js 3.0
Go 1.13

- 構成
- Vue
web配下にvue-cliを用いた一般的な構成としています。
Vuex, VueRouterを用いています。
SPAで完成することを目標としています。

- Go
以下のURLを参考にディレクトリを構成しています。  
https://github.com/golang-standards/project-layout  
機能周りはClean Architectureを参考にディレクトリを構成しています。

root/  
 |  
├cmd/app main.go 配置 ini,logファイルもあります。  
├internal 本プロセスのみで使用するファイルのみ配置。  
 |  　     └domain ドメイン層 model, repositoryを管理  
 |  　     ├infra インフラ層　DB周り  
 |  　     ├interfaces api周り  
 |   　    ├usecase サービス層  
 ├pkg 他プロセスでも共通で使用するであろうファイルを配置 (config, utilなど)  
 ├web フロント周り  

 ## やってきたこと, 苦戦したこと
 - udemyにてGo, Vueの勉強 8月~9月にかけて
 - GoにてAPIサーバーを作る 9月半ば~9月後半

      - サーバーをlocalhostで立ち上げる。 9月半ば
        勉強した内容をもとに作成したので楽に立ち上げる事ができた。

      - ディレクトリ構成を変更 9月半ば

        今回はフレームワークを使わない方針であったため、どのようなディレクトリ構成にするか迷ったが、Go独自のディレクトリ構成でgithub starが多いこともありhttps://github.com/golang-standards/project-layoutを採用

      - Clean Architectureを参考に、各機能のディレクトリ構成を変更 9月半ば

        以下の記事を参考に、testコードが書きやすく保守性が高いシステムにするため。Clean Architectureを実施、DDD TDDは知っていたが、Clean Architectureについては全く知らなかったため、１からキャッチアップ。

        まだ全然わかっていないことが多いので勉強が必要。

        https://christina04.hatenablog.com/entry/go-clean-architecture

        https://engineer.recruit-lifestyle.co.jp/techblog/2018-03-16-go-ddd/

      - mysqlを用いてDB接続 9月後半

        udemyで使用したsqlite3ではなく、使ったことがないmysqlを利用。各機能ごとにクエリを実行するため、クエリ毎にdbConnectionをOpenしてCloseするではなく、サーバーが起動してから停止するまで、Connectionを貼り続けるように実装。これにめちゃくちゃ苦戦する。最終的にエラーの理由がdatabase.goの

        sql.Open()をする時の戻り値を初期化していたことが原因 "="か":="で2日間悩んでいた。現在はデータベース初期化しているファイルとクエリを発行するファイルは同じディレクトリとしている。

      - API機能を実装 9/24

        handlerにてAPI作成する APIを簡素な作りにするため、gorilla/muxライブラリを採用。signup(アカウント登録)の機能を実装する。

        Restletで疎通確認し、mysqlにデータが挿入されることを確認。

        login_idと別にauto_incrementのidを追加。

        エラーハンドリングを追加

      - 画面からsignupができるようにする。9月後半

        フロントからsignupできるようなAPIをaxiosで実装。

        画面からだと通信出来ないことが発覚。→CORSについて勉強

        サーバー側にフロント側からアクセス許可をする実装を追加。  


- ログイン機能を追加 9月後半~10/2

  - サーバー側にログインAPIを実装 9月後半

    signup機能と同様の構成でログイン機能を追加。

    Requestした値をSelectして、DBにデータがあればResponseでOKとする、機能をとりあえず追加。

  - 画面からの疎通確認 9月後半

    signupと同様に疎通しているかを確認

  - 画面にて、アクセス制御を行う。9月後半

    LoginページとSignupページ以外はログインしていなければ遷移出来ないように、VueRouterでアクセス制御を行う。

    beforeEachメソッドをmain.jsに追加

  - 認証機能を追加 9月後半

    jwtを用いたアクセストークン発行機能をGoに追加。

    firebaseに実装を移譲するのではなく(ありきたりなので)、自前でtokenを作成する実装を追加。jwt-goのライブラリを使用。

    signupする時にDBに保存するパスワードを暗号化する実装を追加。 

    ログインするときに、同じ方式で暗号化しDBの暗号化されたパスワードと一致しているか確認し、一致していればアクセストークンを発行。

  - Vuexに処理を移譲 9月後半

    アクセストークンをVuexで管理するために、ログイン機能をstore.jsに移譲。

  - ログイン画面を作成 **現在**

    vuetifyを用いたイケてるログイン画面を作成中、バリデーションも追加する。flexなども対応できるように実装する予定。

    

## TODO

これからアプリ完成に向けて実装する内容を記載。

- ログイン画面、サインアップ画面をイケてる感じにする。スマホも対応。

- pugを使ってhtmlを見やすく整理する。
- go_testを各ファイルに記載する。
- ユーザー固有情報を管理できるようにgo contextパッケージを利用し、並行処理を実現する。
- アカウントロックを実装する。
- メイン画面にチャートを表示する。(apexchart or google chart)
- GAEにデプロイする。(Dockerを使用)
- CircleCIを用いた自動デプロイができるようにする。

- 成行注文ができるようにする。



