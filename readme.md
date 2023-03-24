<samp>

chatgptから返ってくる文章は,GitHubのmarkdownで上手く表示されません。  
適切な位置に空白を挿入し、GitHubのmarkdownで上手く表示されるようにしています。

## usage

```bash
go run main.go
```

標準では、in.mdを読み込み、out.mdに書き出します。  
ファイル名を指定することもできます。  
既に存在するファイルとして出力ファイルが与えられると、そのファイルの末尾に追記されます。

```bash
go run main.go -in before.md -out after.md
```


## example

---

befor: 極座標において、$e^{j\theta}$は、複素数の長さが1で、角度が$\theta$である点を表します。ここで、$j$は虚数単位で、$j^2=-1$です。

---

after: 極座標において、 $e^{j\theta}$ は、複素数の長さが1で、角度が $\theta$ である点を表します。ここで、 $j$ は虚数単位で、 $j^2=-1$ です。

---

</samp>