package components

import (
	. "github.com/pyros2097/gromer/handlebars"
)

var _ = Css(`
	header {
		background: yellow;
	}
`)

func Header() *Template {
	return Html(`
		<header>
			<nav class="navbar" role="navigation" aria-label="main navigation">
				<div class="navbar-brand">
					<a class="navbar-item" href="https://bulma.io">
						<img src="https://bulma.io/images/bulma-logo.png" width="112" height="28">
					</a>

					<a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
						<span aria-hidden="true"></span>
						<span aria-hidden="true"></span>
						<span aria-hidden="true"></span>
					</a>
				</div>

				<div id="navbarBasicExample" class="navbar-menu">
					<div class="navbar-start">
						<a class="navbar-item" href="/">
							Home
						</a>

						<a class="navbar-item" href="/about">
							About
						</a>

						<a class="navbar-item" href="/clock">
							Clock
						</a>

						<a class="navbar-item" href="/counter">
							Counter
						</a>

						<a class="navbar-item" href="/api">
							API
						</a>
					</div>

					<div class="navbar-end">
						<div class="navbar-item">
							<div class="buttons">
								<a class="button is-primary">
									<strong>Sign up</strong>
								</a>
								<a class="button is-light">
									Log in
								</a>
							</div>
						</div>
					</div>
				</div>
			</nav>
		</header>
	`)
}
