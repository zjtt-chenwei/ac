{{define "nav"}}

<nav class="navbar navbar-default">
  <div class="container-fluid">
    <!-- Brand and toggle get grouped for better mobile display -->
    <div class="navbar-header">
      <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
        <span class="sr-only">其他</span>  <!-- 当屏幕小于一定尺寸的时候，导航条会自动隐藏，变成三个“横线”的按钮显示在导航条上，点击这个导航条上的这个按钮，就可以显示出导航的菜单来。-->
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a class="navbar-brand" href="/"></a>
    </div>
   <!-- onmouseover="this.style.color='#f1ad48'" onmouseout="this.style.color='#000000'"  style="color:#000000;" -->
    <!-- Collect the nav links, forms, and other content for toggling -->  
    <div class="navbar-collapse collapse" id="bs-example-navbar-collapse-1">
      <ul class="nav navbar-nav">
        <li><a href="/" class="active">首页 <span class="sr-only">(current)</span></a></li>
        <li><a href="/petmatch" >宠物婚介</a></li>
        <li><a href="../static/baike" >宠物百科</a></li>
        <li><a href="../static/shop" >宠物用品</a></li>
        <li><a href="addpet.html" >关于我们</a></li>
      </ul>
      <!-- <form class="navbar-form navbar-left">
        <div class="form-group">
          <input type="text" class="form-control" placeholder="Search">
        </div>
        <button type="submit" class="btn btn-default">提交</button>
      </form> -->
      <ul class="nav navbar-nav navbar-right">
        {{if $.isLogin}}
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
          {{$.username}}
          <span class="caret"></span>
          </a>
          <ul class="dropdown-menu">
            <li><a href="/people">个人中心</a></li>
            <li><a href="/settings">设置</a></li>
            <li role="separator" class="divider"></li>
            <li><a href="/logout">退出</a></li>
          </ul>
        </li>
        {{else}}
        <li><a href="/login">登录</a></li>
        {{end}}
      </ul>
    </div><!-- /.navbar-collapse -->
  </div><!-- /.container-fluid -->
</nav>

{{end}}