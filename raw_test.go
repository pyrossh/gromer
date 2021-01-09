package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var rawhtml = `
<div>
<p>slug: &ldquo;/blog/gopibot-to-the-rescue&rdquo;
date: &ldquo;2017-10-04&rdquo;</p>

<h2>title: &ldquo;Gopibot To The Rescue&rdquo;</h2>

<p>High Ho Gopibot away!</p>

<p>Everybody please meet Gopibot our chatops bot which I built at Numberz to help us deploy our countless microservices to QA.</p>

<p><img src="../images/posts/gopibot1.png" alt="Gopibot 1" /></p>

<p>So here is the backstory,</p>

<p>I was one of the developers who had access to our QA and Prod servers and the other person was the Head of Engineering
and he is generally a busy guy. So whenever there is a change that needs to be deployed everyone comes to me and tells me
to deploy their microservice/frontend to the QA and blatantly interrupts my awesome coding cycle.</p>

<p>Alright then, I break off from my flow, ssh into the server and start running the deploy command.
And all of you jsdev wannabes who have worked with react and webpack will know the horrors about deploying frontend code right.
It takes forever so I have to wait there looking at the console along with the dev who wanted me to deploy it (lets call him kokill for now).
So kokill and I patiently wait for the webpack build to finish. 1m , 2m, 3m and WTH 15m.
And then its built and the new frontend is deployed to QA. YES! Now I can continue with my work.
But wait then some other dev comes likes call him (D-Ne0) and he asks to deploy something else and again the same process
of ssh’ing the server and another wait. This got repetitive and irritating. Then I started searching for solutions to the problem
and looked high and low and thought that CI/CD is the only thing that can solve this problem. But then I saw something new called ChatOps
where developers have chatbots to talk to automate this manual work. Just like we have bots these days to help you out in your work
like getting your laundry, grocery and making orders.</p>

<p>So I decided to take a shot at this in my free time. And it seems it was simpler than I thought and decided to use Slack our primary
team communication platform. We used it daily for everything and I thought why not have a specific channel just where the bot resides
and people could talk to the bot.</p>

<p>Since we are typically a nodejs shop I decided to find a way to send messages to a slack bot. And slack has this really great sdk for nodejs.
<a href="https://github.com/slackapi/node-slack-sdk">https://github.com/slackapi/node-slack-sdk</a></p>

<p>First I went and created the bot in my slack team settings.
And then wrote a script which would allow it to read messages from the channel it was added.</p>

<p>Here is the simple script,</p>
</div>
`

func TestRawRootTagName(t *testing.T) {
	tests := []struct {
		scenario string
		raw      string
		expected string
	}{
		{
			scenario: "tag set",
			raw: `
			<div>
				<span></span>
			</div>`,
			expected: "div",
		},
		{
			scenario: "tag is empty",
		},
		{
			scenario: "opening tag missing",
			raw:      "</div>",
		},
		{
			scenario: "tag is not set",
			raw:      "div",
		},
		{
			scenario: "tag is not closing",
			raw:      "<div",
		},
		{
			scenario: "tag is not closing",
			raw:      "<div",
		},
		{
			scenario: "tag without value",
			raw:      "<>",
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			tag := rawRootTagName(test.raw)
			require.Equal(t, test.expected, tag)
		})
	}
}

func TestRawMountDismount(t *testing.T) {
	testMountDismount(t, []mountTest{
		{
			scenario: "raw html element",
			node:     Raw(`<h1>Hello</h1>`),
		},
		{
			scenario: "raw svg element",
			node:     Raw(`<svg></svg>`),
		},
	})
}

func TestRawUpdate(t *testing.T) {
	testUpdate(t, []updateTest{
		{
			scenario:   "raw html element returns replace error when updated with a non text-element",
			a:          Raw("<svg></svg>"),
			b:          Div(),
			replaceErr: true,
		},
		{
			scenario: "raw html element is replace by another raw html element",
			a: Div(
				Raw("<div></div>"),
			),
			b: Div(
				Raw("<svg></svg>"),
			),
			matches: []TestUIDescriptor{
				{
					Path:     TestPath(),
					Expected: Div(),
				},
				{
					Path:     TestPath(0),
					Expected: Raw("<svg></svg>"),
				},
			},
		},
		{
			scenario: "raw html element is replace by non-raw html element",
			a: Div(
				Raw("<div></div>"),
			),
			b: Div(
				Text("hello"),
			),
			matches: []TestUIDescriptor{
				{
					Path:     TestPath(),
					Expected: Div(),
				},
				{
					Path:     TestPath(0),
					Expected: Text("hello"),
				},
			},
		},
	})
}
