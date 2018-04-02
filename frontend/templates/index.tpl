<h1>Go demo application</h1>  <a href="/login">Login</a> or <a href="/register">Register</a>  <h2>Usage</h2>
<ol>
    <li>{{if .user.ID}}<span class="text-success">✓</span>{{end}} <a href="/register">Register</a>, Please enter your
        email and password.
    </li>
    <li>{{if .user.ID}}<span class="text-success">✓</span>{{end}} <a href="/login">Login</a>, Enter same email and
        password.
    </li>
</ol>  {{if .user.ID}}
<div class="alert alert-warning">
    <dl>
        <dt>Email</dt>
        <dd>{{.user.Email}}</dd>
        <dt>Password(Hashed)</dt>
        <dd>{{.user.Password}}</dd>
        <dt>API token</dt>
        <dd>{{.user.Token}}</dd>
        <dt>CreatedAt</dt>
        <dd>{{.user.CreatedAt}}</dd>
    </dl>
    <a href="/logout" class="btn btn-primary">Log out</a>
</div>