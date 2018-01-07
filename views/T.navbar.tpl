{{define "navbar"}}

 <div class="container">
 <a href="/" class="navbar-brand">我的博客</a>                
        <ul class="nav navbar-nav">
            <li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
            <li {{if .IsHome}}class="active"{{end}}><a href="/topic">文章</a></li>
        </ul>
    </div>
{{end}}