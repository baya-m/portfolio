# Portfolio

## はじめに
こんにちは。閲覧いただきありがとうございます。

現在転職活動用にポートフォリオを作成しております。
「仮想通貨疑似売買アプリ」を作成しようとしています。

このポートフォリオは現在作成中であり、未完成となっていますので、ご了承ください。

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
 - GoにてAPIサーバーを作る 
 　- サーバーをlocalhostで立ち上げる。 9/12
  勉強した内容をもとに作成したので楽に立ち上げた
  
  
  
  
  
