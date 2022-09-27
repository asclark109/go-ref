-

# Go (Golang)
https://www.youtube.com/watch?v=X4q1OM0voO0&list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N&index=1

Called "Golang" to avoid confusion caused by searching "Go" on search engines.

Features:
- Low Latency (if you are concerned about latency; adopted by twitch)
- Garbage Collection (if you want very fast and lots of support)
- GPU (if concerned about GPU power and utilization)
- Concurrency Support (built-in support)

### Why new language?

Python easy, C++ fast, Java garbage collection. 21st century C programming language. Takes syntax from languages like C, Java and makes it rediculously fast. C,C++ fast but compile time slow. Concurrency was added to these languages later. Golang takes advantage of high memory RAM and GPU power in today's computers.

Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines [golang docs].

### Installation

- `MOD`: for initializing packages...
- `ENV`: environment variables; many env vars set by go. `GoROOT` and `GoWorkSpace` are important. `GoROOT` is where all binaries to run the program are stored. `GoWorkSpace` is the folder where you put your code. Go tells you specifically organize your code in a strict directory structure. Though, this structure is only expected if the code is to go in production or go to a public repository.
- `RUN`:
- `GET`: to get things from the repository
- `FMT`:

Special Directory Structure.
```
project/
  bin/
  pkg/
  src/
    github.com/
      asclark/
        helloWorld  
```
`bin/` holds binaries. `pkg/` holds packages you get from online although that's not compulsory. `src/` where your code goes.

in terminal, see version
```
$ go version
go version go1.19.1 windows/amd64
```

see env variables
```
$ go env
set GO111MODULE=
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\ascla\AppData\Local\go-build
set GOENV=C:\Users\ascla\AppData\Roaming\go\env
set GOEXE=.exe
set GOEXPERIMENT=
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOINSECURE=
set GOMODCACHE=C:\Users\ascla\go\pkg\mod
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=C:\Users\ascla\go
set GOPRIVATE=
set GOPROXY=https://proxy.golang.org,direct
set GOROOT=C:\Program Files\Go
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLDIR=C:\Program Files\Go\pkg\tool\windows_amd64
set GOVCS=
set GOVERSION=go1.19.1
set GCCGO=gccgo
set GOAMD64=v1
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=NUL
set GOWORK=
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -Wl,--no-gc-sections -fmessage-length=0 -fdebug-prefix-map=C:\Users\ascla\AppData\Local\Temp\go-build539226397=/tmp/go-build -gno-record-gcc-switches
```

installation of go located at `set GOROOT=C:\Program Files\Go`. Will help you when running files. GoWorkSpace is indicated by `set GOPATH=C:\Users\ascla\go`









__HTML (Hypertext Markup Language)__ is not a programming language. It is a markup language that tells web browsers how to structure the web pages you visit. It can be as complicated or as simple as the web developer wants it to be. HTML consists of a series of elements, which you use to enclose, wrap, or mark up different parts of content to make it appear or act in a certain way. The enclosing tags can make content into a hyperlink to connect to another page, italicize words, and so on

## Elements, Attributes

### Anatomy of an HTML element

__comments__ are enclosed with __```<! ––```__ and __```––>```__
```html
<!-- This is a comment -->

<!-- Can comment out
 a large number 
 of lines. -->
```

__elements__ are enclosed with __tags__.
An __opening tag__ (__```<tag>```__) marks the start of an element, followed by __content__, ending with a __closing tag__ (__```</tag>```__, demarked by the ```/```).

```html
<p>My cat is very grumpy</p>
```

We can __nest__ elements.

```html
<!-- good -->
<p>My cat is <strong>very</strong> grumpy.</p>
```

The tags have to open and close in a way that they are inside or outside one another. The following does not work

```html
<!-- bad! -->
<p>My cat is <strong>very grumpy.</p></strong>
```

### Block Elements, Inline Elements (different from CSS types of boxes)

Elements can be <u>block elements</u> or <u>inline elements</u>.

__Block-level elements__: form a visible block on a page. A block-level element appears on a new line following the content that precedes it. Any content that follows a block-level element also appears on a new line. Block-level elements are usually structural elements on the page.

>For example, a block-level element might represent headings ```<h>```, paragraphs ```<p>```, lists ```<li>```, navigation menus, or footers. A block-level element wouldn't be nested inside an inline element, but it might be nested inside another block-level element.

__Inline elements__ are contained within block-level elements, and surround only small parts of the document’s content (not entire paragraphs or groupings of content). An inline element will not cause a new line to appear in the document. It is typically used with text.

>For example: an ```<a>``` element creates a hyperlink, and elements such as ```<em>``` or ```<strong>``` create emphasis.

```html
<!-- Inline elements appearing in one line (until overflowing to a new line) -->
<em>first</em><em>second</em><em>third</em>

<!-- each paragraph element (<p></p>) will create a new line-->
<p>fourth</p><p>fifth</p><p>sixth</p>
```

### Empty Elements

Not all elements follow pattern of open tag, content, close tag. Some elements consist of only a single tack (typically used to insert/embed something in the document)

> For example, the ```<img>``` element embeds an image file onto a page:

```html
<img src="https://raw.githubusercontent.com/mdn/beginner-html-site/gh-pages/images/firefox-icon.png">
```

### Attributes

Elements may contain attributes.

```html
<!-- "class" is an example of an attribute-->
<p class="editor-note">My cat is very grumpy</p>
```

__Attributes__: contain extra information about the element that won't appear in the content. The __```class```__ attribute is an identifying name used to target the element with style information.

Attributes should have:
- A space between it and the element name.
    - (For an element with more than one attribute, the attributes should be separated by spaces too.)
- The __attribute name__, followed by an equal sign.
- An __attribute value__, wrapped with opening and closing quote marks.

```html
<p attributeName="attributeValue">
```

Example: The anchor tag /<a> can take a number of attributes.

```html
<p>A link to my <a href="https://www.mozilla.org/" title="The Mozilla homepage" target="_blank">favorite website</a>.</p>
```

Sometimes attributes take boolean values. With these attributes, it is entirely acceptable to write boolean attributes with only the name of the boolean attribute. The value it will take on will be whatever value is inferred by the boolean attribute:
```html
<!-- using the disabled attribute prevents the end user from entering text into the input box -->
<input type="text" disabled="disabled">

<!-- equivalent to -->
<input type="text" disabled>

<!-- text input is allowed, as it doesn't contain the disabled attribute -->
<input type="text">
```

Omitting quotes around attribute values is allowed, but it could break markup in some circumstances (not recommended)

```html
<!-- acceptable -->
<a href=https://www.mozilla.org/>favorite website</a>

<!-- adding title attribute breaks it -->
<!-- interpretted as title attribute, Mozzila attribute, homepage attribute -->
<a href=https://www.mozilla.org/ title=The Mozilla homepage>favorite website</a
```

Single and double quotes allowed
```html
<!-- acceptable -->
<a href="https://www.example.com">A link to my example.</a>
<!-- acceptable -->
<a href='https://www.example.com'>A link to my example.</a>
```

### Anatomy of an HTML document

```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>My test page</title>
  </head>
  <body>
    <p>This is my page</p>
  </body>
</html>
```

1. ```<!DOCTYPE html>``` the __doctype__. doctypes originally were meant to act as links to a set of rules that the HTML page had to follow. More recently, the doctype is a historical artifact that needs to be included for everything else to work right. ```<!DOCTYPE html>``` is the shortest string of characters counting as a valid doc type. Doctypes can also look like:
```html
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
```
2. __```<html></html>```__: The __html element__ /<html> element. This element wraps all the content on the page. It is sometimes known as the <u>root element</u>.
3. __```<head></head>```__: The __head element__. This element acts as a container for everything you want to include on the HTML page, that isn't the content the page will show to viewers. This includes keywords and a page description that would appear in search results, CSS to style content, character set declarations, and more. More on this later.
4. __```<meta charset="utf-8">```__: __Document Character Encoding__ This element specifies the character set for your document to UTF-8, which includes most characters from the vast majority of human written languages. With this setting, the page can now handle any textual content it might contain. There is no reason not to set this, and it can help avoid some problems later.
5. __```<title></title>```__: The __title element__. This sets the title of the page, which is the title that appears in the browser tab the page is loaded in. The page title is also used to describe the page when it is bookmarked.
6. __```<body></body>```__: The __body element__. This contains all the content that displays on the page, including text, images, videos, games, playable audio tracks, or whatever else.

### Whitespace in HTML

The HTML parser reduces all whitespace down to a single space when rendering code.
```html
<!--
    The following below will just appear as two identical lines
    <p> element is a block element, so each <p> element gets its own
    "line"
-->
<p>Dogs are silly.</p>

<p>Dogs        are
         silly.</p>
```

### Special Characters in HTML

In HTML, the following are __special characters__: __```<```__, __```>```__, __```"```__, __```'```__, __```&```__. To have them show in actual HTML content, we reference them by using ```&``` (ampersand) as a special character and then write a code:
```html
<!-- Literal character / Character reference equivalent -->
<	&lt;
>	&gt;
"	&quot;
'	&apos;
&	&amp;
```

## Adding content to ```<head>```

```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>My test page</title>
  </head>
  <body>
    <p>This is my page</p>
  </body>
</html>
```

Note:\
contents added to the ```<head>``` element will not be displayed on the page.\
contents added to the ```<body>``` element will be displayed.
```html
<!-- the head element. Here it is short, but it can get quite full -->
<head>
  <meta charset="utf-8">
  <title>My test page</title>
</head>
```

### Metadata: ```<meta>``` element

__Metadata__ is data that describes data, and HTML has an "official" way of adding metadata to a document — the ```<meta>``` element. Of course, the other stuff we are talking about in this article could also be thought of as metadata too. There are a lot of different types of ```<meta>``` elements that can be included in your page's ```<head>```

```html
<meta charset="utf-8">
```

__```<meta charset>```__ specifies your document character encoding--the character set that the document is permitted to use. utf-8 is a universal character set that includes pretty much any character from any human language. This means that your web page will be able to handle displaying any language; it's therefore a good idea to set this on every web page you create! For example, your page could handle English and Japanese just fine:

As you travel around the web, you'll find other types of metadata, too. A lot of the features you'll see on websites are proprietary creations, designed to provide certain sites (such as social networking sites) with specific pieces of information they can use.

Many <meta> elements include name and content attributes:

__```name```__ specifies the type of meta element it is; what type of information it contains.\
__```content```__ specifies the actual meta content.

```html
<meta name="author" content="Chris Mills">
<meta name="description" content="The MDN Web Docs Learning Area aims to provide
complete beginners to the Web with all they need to know to get
started with developing web sites and applications.">
```

As you travel around the web, you'll find other types of metadata, too. A lot of the features you'll see on websites are proprietary creations, designed to provide certain sites (such as social networking sites) with specific pieces of information they can use.

For example, <a href="https://ogp.me/">Open Graph Data<a> is a metadata protocol that Facebook invented to provide richer metadata for websites. In the MDN Web Docs sourcecode, you'll find this:
```html
<meta property="og:image" content="https://developer.mozilla.org/static/img/opengraph-logo.png">
<meta property="og:description" content="The Mozilla Developer Network (MDN) provides
information about Open Web technologies including HTML, CSS, and APIs for both Web sites
and HTML5 Apps. It also documents Mozilla products, like Firefox OS.">
<meta property="og:title" content="Mozilla Developer Network">
```

One effect of this is that when you link to MDN Web Docs on facebook, the link appears along with an image and description: a richer experience for users.

<p align="center">
    <img src="imgs/facebook-output.png" width="400">
</p>

Twitter also has its own similar proprietary metadata called <a href="https://developer.twitter.com/en/docs/tweets/optimize-with-cards/overview/abouts-cards">Twitter Cards<a>, which has a similar effect when the site's URL is displayed on twitter.com. For example:

```html
<meta name="twitter:title" content="Mozilla Developer Network">
```

### Adding custom icons to a site

To further enrich your site design, you can add references to custom icons in your metadata, and these will be displayed in certain contexts. The most commonly used of these is the __favicon__ (short for "favorites icon", referring to its use in the "favorites" or "bookmarks" lists in browsers).

The humble favicon has been around for many years. It is the first icon of this type: a 16-pixel square icon used in multiple places. You may see (depending on the browser) favicons displayed in the browser tab containing each open page, and next to bookmarked pages in the bookmarks panel.

A favicon can be added to your page by:

1. Saving it in the same directory as the site's index page, saved in __```.ico```__ format (most browsers will support favicons in more common formats like __```.gif```__ or __```.png```__, but using the ICO format will ensure it works as far back as Internet Explorer 6.)
2. Adding the following line into your HTML's ```<head>``` block to reference it:
```html
<link rel="icon" href="favicon.ico" type="image/x-icon">
```

There are lots of other icon types to consider these days as well. For example, you'll find this in the source code of the MDN Web Docs homepage:

```html
<!-- third-generation iPad with high-resolution Retina display: -->
<link rel="apple-touch-icon-precomposed" sizes="144x144" href="https://developer.mozilla.org/static/img/favicon144.png">
<!-- iPhone with high-resolution Retina display: -->
<link rel="apple-touch-icon-precomposed" sizes="114x114" href="https://developer.mozilla.org/static/img/favicon114.png">
<!-- first- and second-generation iPad: -->
<link rel="apple-touch-icon-precomposed" sizes="72x72" href="https://developer.mozilla.org/static/img/favicon72.png">
<!-- non-Retina iPhone, iPod Touch, and Android 2.1+ devices: -->
<link rel="apple-touch-icon-precomposed" href="https://developer.mozilla.org/static/img/favicon57.png">
<!-- basic favicon -->
<link rel="icon" href="https://developer.mozilla.org/static/img/favicon32.png">
```

### Applying CSS, Javascript to HTML

Just about all websites you'll use in the modern day will employ CSS to make them look cool, and JavaScript to power interactive functionality, such as video players, maps, games, and more. These are most commonly applied to a web page using the ```<link>``` element and the ```<script>``` element, respectively.

The ```<link>``` element should always go inside the head of your document. This takes two attributes, ```rel="stylesheet"```, which indicates that it is the document's stylesheet, and ```href```, which contains the path to the stylesheet file:

```html
<link rel="stylesheet" href="my-css-file.css">
```

The ```<script>``` element should also go into the head, and should include a src attribute containing the path to the JavaScript you want to load, and defer, which basically instructs the browser to load the JavaScript after the page has finished parsing the HTML. This is useful as it makes sure that the HTML is all loaded before the JavaScript runs, so that you don't get errors resulting from JavaScript trying to access an HTML element that doesn't exist on the page yet. There are actually a number of ways to handle loading JavaScript on your page, but this is the most foolproof one to use for modern browsers (for others, read Script loading strategies).

```html
<script src="my-js-file.js" defer></script>

<!-- The <script> element may look like an empty element, but it's not, and so needs a closing tag. Instead of pointing to an external script file, you can also choose to put your script inside the <script> element-->
```

## HTML text fundamentals

### headings ```<h1>```, paragraphs ```<p>```

Structured content makes the reading experience easier and more enjoyable. 
In HTML, each paragraph has to be wrapped in a ```<p>``` element, like so:

```html
<p>I am a paragraph, oh yes I am.</p>
```

Each heading has to be wrapped in a heading element:
```html
<h1>I am the title of the story.</h1>
```

There are six heading elements: ```<h1>```, ```<h2>```, ```<h3>```, ```<h4>```, ```<h5>```, and ```<h6>```. Each element represents a different level of content in the document; ```<h1>``` represents the main heading, ```<h2>``` represents subheadings, ```<h3>``` represents sub-subheadings, and so on.

```html
<!-- Example structured text -->
<h1>The Crushing Bore</h1>

<p>By Chris Mills</p>

<h2>Chapter 1: The dark night</h2>

<p>It was a dark night. Somewhere, an owl hooted. The rain lashed down on the ...</p>

<h2>Chapter 2: The eternal silence</h2>

<p>Our protagonist could not so much as a whisper out of the shadowy figure ...</p>

<h3>The specter speaks</h3>

<p>Several more hours had passed, when all of a sudden the specter sat bolt upright and exclaimed, "Please have mercy on my soul!"</p>
```

It's really up to you what the elements involved represent, as long as the hierarchy makes sense. You just need to bear in mind a few best practices as you create such structures:

Preferably, <u>you should use a single ```<h1>``` per page</u>—this is the top level heading, and all others sit below this in the hierarchy.

Make sure you use the headings in the correct order in the hierarchy. Don't use ```<h3>``` elements to represent subheadings, followed by ```<h2>``` elements to represent sub-subheadings—that doesn't make sense and will lead to weird results.

Of the six heading levels available, you should aim to use no more than three per page, unless you feel it is necessary. Documents with many levels (for example, a deep heading hierarchy) become unwieldy and difficult to navigate. On such occasions, it is advisable to spread the content over multiple pages if possible.

<u>Structural markup is useful because to style content with CSS, or make it do interesting things with JavaScript, you need to have elements wrapping the relevant content, so CSS/JavaScript can effectively target it.</u>

### Semantics

we need to make sure we are using the correct elements, giving our content the correct meaning, function, or appearance. In this context, the ```<h1>``` element is also a __semantic element__, which gives the text it wraps around the role (or meaning) of "a top level heading on your page."

```html
<h1>This is a top level heading</h1>
```

By default, the browser will give it a large font size to make it look like a heading (although you could style it to look like anything you wanted using CSS). More importantly, its semantic value will be used in multiple ways, for example by search engines and screen readers (as mentioned above).

On the other hand, you could make any element look like a top level heading. Consider the following:

```html
<span style="font-size: 32px; margin: 21px 0; display: block;">Is this a top level heading?</span>
```

This is a __```<span>```__ element. It has no semantics. You use it to wrap content when you want to apply CSS to it (or do something to it with JavaScript) without giving it any extra meaning. (You'll find out more about these later on in the course.) We've applied some CSS to it to make it look like a top level heading, but since it has no semantic value, it will not get any of the extra benefits described above. It is a good idea to use the relevant HTML element for the job.

### Lists

__```<ul>```__: __unordered list__

__```<li>```__: __list item__
```html
<ul>
  <li>milk</li>
  <li>eggs</li>
  <li>bread</li>
  <li>hummus</li>
</ul>
```

__```<ol>```__: __ordered list__
<ol>
  <li>Drive to the end of the road</li>
  <li>Turn right</li>
  <li>Go straight across the first two roundabouts</li>
  <li>Turn left at the third roundabout</li>
  <li>The school is on your right, 300 meters up the road</li>
</ol>

can <u>nest</u> lists
```html
<ol>
  <li>Remove the skin from the garlic, and chop coarsely.</li>
  <li>Remove all the seeds and stalk from the pepper, and chop coarsely.</li>
  <li>Add all the ingredients into a food processor.</li>
  <li>Process all the ingredients into a paste.
    <ul>
      <li>If you want a coarse "chunky" hummus, process it for a short time.</li>
      <li>If you want a smooth hummus, process it for a longer time.</li>
    </ul>
  </li>
</ol>
```

### Emphasis

When we want to add emphasis in spoken language, we stress certain words, subtly altering the meaning of what we are saying. Similarly, in written language we tend to stress words by putting them in italics.

__```<em>```__ (emphasis): makes text italicized

In HTML we use the ```<em>``` (emphasis) element to mark up such instances. As well as making the document more interesting to read, these are recognized by screen readers and spoken out in a different tone of voice. Browsers style this as italic by default, but <u>you shouldn't use this tag purely to get italic styling</u>. To do that, you'd use a ```<span>``` element and some CSS, or perhaps an ```<i>``` element (see below).

```html
<p>I am <em>glad</em> you weren't <em>late</em>.</p>
```

__```<strong>```__ (bold): makes text bold

In HTML we use the ```<strong>``` (strong importance) element to mark up such instances. As well as making the document more useful, again these are recognized by screen readers and spoken in a different tone of voice. Browsers style this as bold text by default, but you shouldn't use this tag purely to get bold styling. To do that, you'd use a ```<span>``` element and some CSS, or perhaps a ```<b>``` element (see below).

```html
<p>This liquid is <strong>highly toxic</strong>.</p>

<p>I am counting on you. <strong>Do not</strong> be late!</p>
```

Can nest.
```html
<p>This liquid is <strong>highly toxic</strong> —
if you drink it, <strong>you may <em>die</em></strong>.</p>
```

### continued

The elements we've discussed so far have clearcut associated semantics. The situation with ```<b>```, ```<i>```, and ```<u>``` is somewhat more complicated. They came about so people could write bold, italics, or underlined text in an era when CSS was still supported poorly or not at all. Elements like this, which only affect presentation and not semantics, are known as __presentational elements__ and should no longer be used because, as we've seen before, semantics is so important to accessibility, SEO, etc.

HTML5 redefined ```<b>```, ```<i>```, and ```<u>``` with new, somewhat confusing, semantic roles.

Here's the best rule of thumb: It's likely appropriate to use ```<b>```, ```<i>```, or ```<u>``` to convey a meaning traditionally conveyed with bold, italics, or underline, provided there is no more suitable element. However, it always remains critical to keep an accessibility mindset. The concept of italics isn't very helpful to people using screen readers, or to people using a writing system other than the Latin alphabet.

__```<i>```__ is used to convey a meaning traditionally conveyed by italic: foreign words, taxonomic designation, technical terms, a thought...\
__```<b>```__ is used to convey a meaning traditionally conveyed by bold: key words, product names, lead sentence...\
__```<u>```__ is used to convey a meaning traditionally conveyed by underline: proper name, misspelling.

> Note: People strongly associate underlining with hyperlinks. Therefore, on the web, it's best to underline only links. Use the ```<u>``` element when it's semantically appropriate, but consider using CSS to change the default underline to something more appropriate on the web. The example below illustrates how it can be done.

```html
<!-- scientific names -->
<p>
  The Ruby-throated Hummingbird (<i>Archilochus colubris</i>)
  is the most common hummingbird in Eastern North America.
</p>

<!-- foreign words -->
<p>
  The menu was a sea of exotic words like <i lang="uk-latn">vatrushka</i>,
  <i lang="id">nasi goreng</i> and <i lang="fr">soupe à l'oignon</i>.
</p>

<!-- a known misspelling -->
<p>
  Someday I'll learn how to <u style="text-decoration-line: underline; text-decoration-style: wavy;">spel</u> better.
</p>

<!-- Highlight keywords in a set of instructions -->
<ol>
  <li>
    <b>Slice</b> two pieces of bread off the loaf.
  </li>
  <li>
    <b>Insert</b> a tomato slice and a leaf of
    lettuce between the slices of bread.
  </li>
</ol>
```

## Creating Hyperlinks

Hyperlinks are one of the most exciting innovations the Web has to offer. They've been a feature of the Web since the beginning, and are what makes the Web a web. Hyperlinks allow us to link documents to other documents or resources, link to specific parts of documents, or make apps available at a web address. Almost any web content can be converted to a link so that when clicked or otherwise activated the web browser goes to another web address

__Internet__: is a worldwide network of networks that uses the Internet protocol suite (also named TCP/IP from its two most important protocols).

__HyperText Transfer Protocol (HTTP)__: underlying network protocol that enables transfer of hypermedia documents on the Web, typically between a browser and a server so that humans can read them. The current version of the HTTP specification is called HTTP/2.

__Uniform Resource Loader (URL)__: is a text string that specifies where a resource (such as a web page, image, or video) can be found on the Internet.

In the context of HTTP, URLs are called "Web address" or "link". Your browser displays URLs in its address bar, for example: https://developer.mozilla.org Some browsers display only the part of a URL after the "//", that is, the Domain name.

__Domain name__: is a website's address on the Internet. Domain names are used in URLs to identify to which server belong a specific webpage. The domain name consists of a hierarchical sequence of names (labels) separated by periods (dots) and ending with an extension

URLs can also be used for file transfer (FTP) , emails (SMTP), and other applications.

> Note: A URL can point to HTML files, text files, images, text documents, video and audio files, or anything else that lives on the Web. If the web browser doesn't know how to display or handle the file, it will ask you if you want to open the file (in which case the duty of opening or handling the file is passed to a suitable native app on the device) or download the file (in which case you can try to deal with it later on).

### Anatomy of a link

A basic link is created by wrapping the text or other content, see Block level links, inside an ```<a>``` element and using the ```href``` attribute, also known as a Hypertext Reference, or target, that contains the web address.

```html
<p>I'm creating a link to
<a href="https://www.mozilla.org/en-US/">the Mozilla homepage</a>.
</p>
```

Another attribute you may want to add to your links is ```title```. The title contains additional information about the link, such as which kind of information the page contains, or things to be aware of on the web site.
```html
<p>I'm creating a link to
<a href="https://www.mozilla.org/en-US/"
   title="The best place to find more information about Mozilla's
          mission and how to contribute">the Mozilla homepage</a>.
</p>
```

### Block level links

As mentioned before, almost any content can be made into a link, even <i>block-level elements</i>. If you have an image you want to make into a link, use the ```<a>``` element and reference the image file with the ```<img>``` element.
```html
<a href="https://www.mozilla.org/en-US/">
  <img src="mozilla-image.png" alt="mozilla logo that links to the mozilla homepage">
</a>
```

### Quick aside on URLs and Paths

To fully understand link targets, you need to understand URLs and file paths. This section gives you the information you need to achieve this.

A URL, or Uniform Resource Locator is a string of text that defines where something is located on the Web. For example, Mozilla's English homepage is located at https://www.mozilla.org/en-US/.

URLs use paths to find files. Paths specify where the file you're interested in is located in the filesystem.

Let's do an example.
<p align="center">
    <img src="imgs/simple-directory.png" width="700">
</p>

The __root__ of this directory structure is called ```creating-hyperlinks```. When working locally with a web site, you'll have one directory that contains the entire site. Inside the root, we have an ```index.html``` file and a ```contacts.html```. In a real website, ```index.html``` would be our home page or landing page (a web page that serves as the entry point for a website or a particular section of a website.).

There are also two directories inside our root — ```pdfs``` and ```projects```. These each have a single file inside them — a PDF (```project-brief.pdf```) and an ```index.html``` file, respectively. <u>Note that you can have two ```index.html``` files in one project</u>, as long as they're in different filesystem locations. The second ```index.html``` would perhaps be the main landing page for project-related information.

__Same directory__: If you wanted to include a hyperlink inside ```index.html``` (the top level index.html) pointing to ```contacts.html```, you would specify the filename that you want to link to, because it's in the same directory as the current file. The URL you would use is ```contacts.html```:
```html
<p>Want to contact a specific staff member?
Find details on our <a href="contacts.html">contacts page</a>.</p>
```

__Moving down into subdirectories__: If you wanted to include a hyperlink inside ```index.html``` (the top level index.html) pointing to ```projects/index.html```, you would need to go down into the projects directory before indicating the file you want to link to. This is done by specifying the directory's name, then a forward slash, then the name of the file. The URL you would use is ```projects/index.html```:
```html
<p>Visit my <a href="projects/index.html">project homepage</a>.</p>
```

__Moving back up into parent directories__: If you wanted to include a hyperlink inside ```projects/index.html``` pointing to ```pdfs/project-brief.pdf```, you'd have to go up a directory level, then back down into the ```pdf``` directory. To go up a directory, use two dots — ```..``` — so the URL you would use is ```../pdfs/project-brief.pdf```:
```html
<p>A link to my <a href="../pdfs/project-brief.pdf">project brief</a>.</p>
```

> Note: You can combine multiple instances of these features into complex URLs, if needed, for example: ```../../../complex/path/to/my/file.html```.

### Linking to document fragments

It's possible to link to a specific part of an HTML document, known as a __document fragment__, rather than just to the top of the document. To do this you first have to assign an __```id```__ attribute to the element you want to link to. It normally makes sense to link to a specific heading, so this would look something like the following:
```html
<h2 id="Mailing_address">Mailing address</h2>
```

Then to link to that specific id, you'd include it at the end of the URL, preceded by a hash/pound symbol (```#```), for example:

```html
<p>Want to write us a letter? Use our <a href="contacts.html#Mailing_address">mailing address</a>.</p>
```

You can even use the document fragment reference on its own to link to another part of the current document:

```html
<p>The <a href="#Mailing_address">company mailing address</a> can be found at the bottom of this page.</p>
```

### Absolute, Relative Paths

__absolute URL__: Points to a location defined by its absolute location on the web, including protocol and domain name. 

> if an ```index.html``` page is uploaded to a directory called ```projects``` that sits inside the __root__ of a web server, and the web site's __domain__ is ```https://www.example.com```, the page would be available at ```https://www.example.com/projects/index.html``` (or even just ```https://www.example.com/projects/```, as <u>most web servers just look for a landing page such as index.html to load if it isn't specified in the URL</u>.

An absolute URL will always point to the same location, no matter where it's used.

__relative URL__: Points to a location that is relative to the file you are linking from, more like what we looked at in the previous section. 

> if we wanted to link from our example file at ```https://www.example.com/projects/index.html``` to a PDF file in the same directory, the URL would just be the filename — ```project-brief.pdf``` — no extra information needed. If the PDF was available in a subdirectory inside ```projects``` called ```pdfs```, the relative link would be ```pdfs/project-brief.pdf``` (the equivalent absolute URL would be ```https://www.example.com/projects/pdfs/project-brief.pdf```.)

A relative URL will point to different places depending on the actual location of the file you refer from — for example if we moved our ```index.html``` file out of the ```projects``` directory and into the __root__ of the web site (the top level, not in any directories), the ```pdfs/project-brief.pdf``` relative URL link inside it would now point to a file located at ```https://www.example.com/pdfs/project-brief.pdf```, not a file located at ```https://www.example.com/projects/pdfs/project-brief.pdf```.

Of course, the location of the ```project-brief.pdf``` file and ```pdfs``` folder won't suddenly change because you moved the ```index.html``` file — this would make your link point to the wrong place, so it wouldn't work if clicked on. You need to be careful!

## Link best practices

### Use clear link wording.

- Screenreader users like jumping around from link to link on the page, and reading links out of context.
- Search engines use link text to index target files, so it is a good idea to include keywords in your link text to effectively describe what is being linked to.
- Visual readers skim over the page rather than reading every word, and their eyes will be drawn to page features that stand out, like links. They will find descriptive link text useful.

```html
<!-- good -->
<p><a href="https://firefox.com/">
    Download Firefox
</a></p>

<!-- bad -->
<p><a href="https://firefox.com/">
  Click here
</a>
to download Firefox</p>
```

<u>Tips</u>

- __Don't repeat the URL as part of the link text__ — URLs look ugly, and sound even uglier when a screen reader reads them out letter by letter.
- __Don't say "link" or "links to" in the link text__ — it's just noise. Screen readers tell people there's a link. Visual users will also know there's a link, because links are generally styled in a different color and underlined (this convention generally shouldn't be broken, as users are used to it).
- __Keep your link text as short as possible__ — this is helpful because screen readers need to interpret the entire link text.
- __Minimize instances where multiple copies of the same text are linked to different places__. This can cause problems for screen reader users, if there's a list of links out of context that are labeled "click here", "click here", "click here".

### Linking to non-HTML resources — leave clear signposts

When linking to a resource that will be downloaded (like a PDF or Word document), streamed (like video or audio), or has another potentially unexpected effect (opens a popup window, or loads a Flash movie), you should add clear wording to reduce any confusion.

For example:

> If you're on a low bandwidth connection, click a link, and then a multiple megabyte download starts unexpectedly.

> If you don't have the Flash player installed, click a link, and then suddenly get taken to a page that requires Flash.

```html
<!-- good examples -->
<p><a href="https://www.example.com/large-report.pdf">
  Download the sales report (PDF, 10MB)
</a></p>

<p><a href="https://www.example.com/video-stream/" target="_blank">
  Watch the video (stream opens in separate tab, HD quality)
</a></p>

<p><a href="https://www.example.com/car-game">
  Play the car game (requires Flash)
</a></p>
```

<u>When you are linking to a resource that's to be downloaded rather than opened in the browser</u>, you can use the __```download```__ attribute to provide a default save filename.
```html
<a href="https://download.mozilla.org/?product=firefox-latest-ssl&os=win64&lang=en-US"
   download="firefox-latest-64bit-installer.exe">
  Download Latest Firefox for Windows (64-bit) (English, US)
</a>
```

It's possible to create links or buttons that, when clicked, <u>open a new outgoing email message rather than linking to a resource or page</u>. This is done using the ```<a>``` element and the __```mailto:```__ URL scheme.

In its most basic and commonly used form, a ```mailto:``` link indicates the email address of the intended recipient. 
```html
<a href="mailto:nowhere@mozilla.org">Send email to nowhere</a>
```

## Advanced Text Formatting

### Description Lists

In HTML text fundamentals, we walked through how to mark up basic lists in HTML, but we didn't mention the third type of list you'll occasionally come across — __description lists__. The purpose of these lists is <u>to mark up a set of items and their associated descriptions</u>, such as terms and definitions, or questions and answers. 

```html
<dl>
  <dt>soliloquy</dt>
  <dd>In drama, where a character speaks to themselves, representing their inner thoughts or feelings and in the process relaying them to the audience (but not to other characters.)</dd>
  <dt>monologue</dt>
  <dd>In drama, where a character speaks their thoughts out loud to share them with the audience and any other characters present.</dd>
  <dt>aside</dt>
  <dd>In drama, where a character shares a comment only with the audience for humorous or dramatic effect. This is usually a feeling, thought, or piece of additional background information.</dd>
</dl>
```

Description lists use a different wrapper than the other list types — __```<dl>```__; in addition each term is wrapped in a <dt> (description term) element, and each description is wrapped in a <dd> (description definition) element.

__```<dl>```__: __description list__
__```<dt>```__: __description term__
__```<dd>```__: __description definition__

<u> can have single term with multiple decsriptions </u>

### Quotations

__```<blockquote>```__: __Blockquotes__, block level content quotation

If a section of block level content (be it a paragraph, multiple paragraphs, a list, etc.) is quoted from somewhere else, you should wrap it inside a ```<blockquote>``` element to signify this, and include a URL pointing to the source of the quote inside a cite attribute. For example, the following markup is taken from the MDN ```<blockquote>``` element page:

```html
<p>Here below is a blockquote...</p>
<blockquote cite="https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote">
  <p>The <strong>HTML <code>&lt;blockquote&gt;</code> Element</strong> (or <em>HTML Block
  Quotation Element</em>) indicates that the enclosed text is an extended quotation.</p>
</blockquote>
```

__```<q>```__: __inline quote__
```html
<p>The quote element — <code>&lt;q&gt;</code> — is <q cite="https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q">intended
for short quotations that don't require paragraph breaks.</q></p>
```

### Citations

The content of the ```cite``` attribute sounds useful, but unfortunately browsers, screenreaders, etc. don't really do much with it. There is no way to get the browser to display the contents of cite, without writing your own solution using JavaScript or CSS. If you want to make the source of the quotation available on the page you need to make it available in the text via a link or some other appropriate way.

There is a ```<cite>``` element, but this is meant to contain the title of the resource being quoted, e.g. the name of the book. There is no reason, however, why you couldn't link the text inside ```<cite>``` to the quote source in some way:
```html
<p>According to the <a href="/en-US/docs/Web/HTML/Element/blockquote">
<cite>MDN blockquote page</cite></a>:
</p>

<blockquote cite="https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote">
  <p>The <strong>HTML <code>&lt;blockquote&gt;</code> Element</strong> (or <em>HTML Block
  Quotation Element</em>) indicates that the enclosed text is an extended quotation.</p>
</blockquote>

<p>The quote element — <code>&lt;q&gt;</code> — is <q cite="https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q">intended
for short quotations that don't require paragraph breaks.</q> -- <a href="/en-US/docs/Web/HTML/Element/q">
<cite>MDN q page</cite></a>.</p>
```

### Abbreviations

Another fairly common element you'll meet when looking around the Web is __```<abbr>```__ — this is used to wrap around an abbreviation or acronym, and provide a full expansion of the term (included inside a title attribute.)

```html
<!-- When user hovers over "HTML", "Hypertext Markup Language" appears-->
<p>We use <abbr title="Hypertext Markup Language">HTML</abbr> to structure our web documents.</p>

<p>I think <abbr title="Reverend">Rev.</abbr> Green did it in the kitchen with the chainsaw.</p>
```

### Marking up contact details

HTML has an element for marking up contact details — __```<address>```__. This wraps around your contact details

```html
<address>
  Chris Mills, Manchester, The Grim North, UK
</address>
```

It could also include more complex markup, and other forms of contact information

```html
<address>
  <p>
    Chris Mills<br>
    Manchester<br>
    The Grim North<br>
    UK
  </p>

  <ul>
    <li>Tel: 01234 567 890</li>
    <li>Email: me@grim-north.co.uk</li>
  </ul>
</address>
```

Note that something like this would also be OK, if the linked page contained the contact information:

```html
<address>
  Page written by <a href="../authors/chris-mills/">Chris Mills</a>.
</address>
```

### Subscript, superscript

You will occasionally need to use superscript and subscript when marking up items like dates, chemical formulae, and mathematical equations so they have the correct meaning. The ```<sup>``` and ```<sub>``` elements handle this job. 

```html
<p>My birthday is on the 25<sup>th</sup> of May 2001.</p>
<p>Caffeine's chemical formula is C<sub>8</sub>H<sub>10</sub>N<sub>4</sub>O<sub>2</sub>.</p>
<p>If x<sup>2</sup> is 9, x must equal 3 or -3.</p>
```

### Representing Computer Code

There are a number of elements available for marking up computer code using HTML:

__```<code>```__: For marking up generic pieces of computer code.\
__```<pre>```__: For retaining whitespace (generally code blocks) — if you use indentation or excess whitespace inside your text, browsers will ignore it and you will not see it on your rendered page. If you wrap the text in __```<pre></pre>```__ tags however, your whitespace will be rendered identically to how you see it in your text editor.\
__```<var>```__: For specifically marking up variable names.
__```<kbd>```__: For marking up keyboard (and other types of) input entered into the computer.
__```<samp>```__: For marking up the output of a computer program.

### Marking up times and dates

HTML also provides the __```<time>```__ element for marking up times and dates in a machine-readable format. 

```html
<!-- Standard simple date -->
<time datetime="2016-01-20">20 January 2016</time>
<!-- Just year and month -->
<time datetime="2016-01">January 2016</time>
<!-- Just month and day -->
<time datetime="01-20">20 January</time>
<!-- Just time, hours and minutes -->
<time datetime="19:30">19:30</time>
<!-- You can do seconds and milliseconds too! -->
<time datetime="19:30:01.856">19:30:01.856</time>
<!-- Date and time -->
<time datetime="2016-01-20T19:30">7.30pm, 20 January 2016</time>
<!-- Date and time with timezone offset -->
<time datetime="2016-01-20T19:30+01:00">7.30pm, 20 January 2016 is 8.30pm in France</time>
<!-- Calling out a specific week number -->
<time datetime="2016-W04">The fourth week of 2016</time>
```

## Document, Website Structure

In addition to defining individual parts of your page (such as "a paragraph" or "an image"), HTML also boasts a number of block level elements used to define areas of your website (such as "the header", "the navigation menu", "the main content column"). This article looks into how to plan a basic website structure, and write the HTML to represent this structure.

__header__: Usually a big strip across the top with a big heading, logo, and perhaps a tagline. This usually stays the same from one webpage to another.

__navigation bar__: Links to the site's main sections; usually represented by menu buttons, links, or tabs. Like the header, this content usually remains consistent from one webpage to another — having inconsistent navigation on your website will just lead to confused, frustrated users. Many web designers consider the navigation bar to be part of the header rather than an individual component, but that's not a requirement; in fact, some also argue that having the two separate is better for accessibility, as screen readers can read the two features better if they are separate.

__main content__: A big area in the center that contains most of the unique content of a given webpage, for example, the video you want to watch, or the main story you're reading, or the map you want to view, or the news headlines, etc. This is the one part of the website that definitely will vary from page to page!

__sidebar__: Some peripheral info, links, quotes, ads, etc. Usually, this is contextual to what is contained in the main content (for example on a news article page, the sidebar might contain the author's bio, or links to related articles) but there are also cases where you'll find some recurring elements like a secondary navigation system.

__footer__: A strip across the bottom of the page that generally contains fine print, copyright notices, or contact info. It's a place to put common information (like the header) but usually, that information is not critical or secondary to the website itself. The footer is also sometimes used for SEO purposes, by providing links for quick access to popular content.

In your HTML code, you can mark up sections of content based on their functionality — you can use elements that represent the sections of content described above unambiguously, and assistive technologies like screenreaders can recognize those elements and help with tasks like "find the main navigation", or "find the main content." 

To implement such semantic mark up, HTML provides dedicated tags that you can use to represent such sections, for example:

header: __```<header>```__.
navigation bar: __```<nav>```__.
main content: __```<main>```__, with various content subsections represented by __```<article>```__, __```<section>```__, and __```<div>```__ elements.
sidebar: __```<aside>```__; often placed inside __```<main>```__.
footer: __```<footer>```__.

Example:
```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">

    <title>My page title</title>
    <link href="https://fonts.googleapis.com/css?family=Open+Sans+Condensed:300|Sonsie+One" rel="stylesheet" type="text/css">
    <link rel="stylesheet" href="style.css">

    <!-- the below three lines are a fix to get HTML5 semantic elements working in old versions of Internet Explorer-->
    <!--[if lt IE 9]>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.3/html5shiv.js"></script>
    <![endif]-->
  </head>

  <body>
    <!-- Here is our main header that is used across all the pages of our website -->

    <header>
      <h1>Header</h1>
    </header>

    <nav>
      <ul>
        <li><a href="#">Home</a></li>
        <li><a href="#">Our team</a></li>
        <li><a href="#">Projects</a></li>
        <li><a href="#">Contact</a></li>
      </ul>

       <!-- A Search form is another common non-linear way to navigate through a website. -->

       <form>
         <input type="search" name="q" placeholder="Search query">
         <input type="submit" value="Go!">
       </form>
     </nav>

    <!-- Here is our page's main content -->
    <main>

      <!-- It contains an article -->
      <article>
        <h2>Article heading</h2>

        <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Donec a diam lectus. Set sit amet ipsum mauris. Maecenas congue ligula as quam viverra nec consectetur ant hendrerit. Donec et mollis dolor. Praesent et diam eget libero egestas mattis sit amet vitae augue. Nam tincidunt congue enim, ut porta lorem lacinia consectetur.</p>

        <h3>Subsection</h3>

        <p>Donec ut librero sed accu vehicula ultricies a non tortor. Lorem ipsum dolor sit amet, consectetur adipisicing elit. Aenean ut gravida lorem. Ut turpis felis, pulvinar a semper sed, adipiscing id dolor.</p>

        <p>Pelientesque auctor nisi id magna consequat sagittis. Curabitur dapibus, enim sit amet elit pharetra tincidunt feugiat nist imperdiet. Ut convallis libero in urna ultrices accumsan. Donec sed odio eros.</p>

        <h3>Another subsection</h3>

        <p>Donec viverra mi quis quam pulvinar at malesuada arcu rhoncus. Cum soclis natoque penatibus et manis dis parturient montes, nascetur ridiculus mus. In rutrum accumsan ultricies. Mauris vitae nisi at sem facilisis semper ac in est.</p>

        <p>Vivamus fermentum semper porta. Nunc diam velit, adipscing ut tristique vitae sagittis vel odio. Maecenas convallis ullamcorper ultricied. Curabitur ornare, ligula semper consectetur sagittis, nisi diam iaculis velit, is fringille sem nunc vet mi.</p>
      </article>

      <!-- the aside content can also be nested within the main content -->
      <aside>
        <h2>Related</h2>

        <ul>
          <li><a href="#">Oh I do like to be beside the seaside</a></li>
          <li><a href="#">Oh I do like to be beside the sea</a></li>
          <li><a href="#">Although in the North of England</a></li>
          <li><a href="#">It never stops raining</a></li>
          <li><a href="#">Oh well...</a></li>
        </ul>
      </aside>

    </main>

    <!-- And here is our main footer that is used across all the pages of our website -->

    <footer>
      <p>©Copyright 2050 by nobody. All rights reversed.</p>
    </footer>

  </body>
</html>
```

<p align="center">
    <img src="imgs/sample-website.png" width="800">
</p>

It's good to understand the overall meaning of all the HTML sectioning elements in detail — this is something you'll work on gradually as you start to get more experience with web development. You can find a lot of detail by reading our HTML element reference. For now, these are the main definitions that you should try to understand:

__```<main>```__ is for content unique to this page. Use ```<main>``` only once per page, and put it directly inside ```<body>```. Ideally this shouldn't be nested within other elements.

__```<article>```__ encloses a block of related content that makes sense on its own without the rest of the page (e.g., a single blog post).

__```<section>```__ is similar to ```<article>```, but it is more for grouping together a single part of the page that constitutes one single piece of functionality (e.g., a mini map, or a set of article headlines and summaries), or a theme. It's considered best practice to begin each section with a heading; also note that you can break ```<article>```s up into different ```<section>```s, or ```<section>```s up into different ```<article>```s, depending on the context.

__```<aside>```__ contains content that is not directly related to the main content but can provide additional information indirectly related to it (glossary entries, author biography, related links, etc.).

__```<header>```__ represents a group of introductory content. If it is a child of ```<body>``` it defines the global header of a webpage, but if it's a child of an ```<article>``` or ```<section>``` it defines a specific header for that section (try not to confuse this with titles and headings).

__```<nav>```__ contains the main navigation functionality for the page. Secondary links, etc., would not go in the navigation.

__```<footer>```__ represents a group of end content for a page.

### Non-semantic Wrappers

Sometimes you'll come across a situation where you can't find an ideal semantic element to group some items together or wrap some content. Sometimes you might want to just group a set of elements together to affect them all as a single entity with some CSS or JavaScript. For cases like these, HTML provides the ```<div>``` and ```<span>``` elements. You should use these preferably with a suitable class attribute, to provide some kind of label for them so they can be easily targeted.

__```<span>```__: is an inline non-semantic element, which you should only use if you can't think of a better semantic text element to wrap your content, or don't want to add any specific meaning.
```html
<p>The King walked drunkenly back to his room at 01:00, the beer doing nothing to aid
him as he staggered through the door <span class="editor-note">[Editor's note: At this point in the
play, the lights should be down low]</span>.</p>
```

In this case, the editor's note is supposed to merely provide extra direction for the director of the play; it is not supposed to have extra semantic meaning. For sighted users, CSS would perhaps be used to distance the note slightly from the main text.

__```<div>```__: is a block level non-semantic element, which you should only use if you can't think of a better semantic block element to use, or don't want to add any specific meaning. For example, imagine a shopping cart widget that you could choose to pull up at any point during your time on an e-commerce site:
```html
<div class="shopping-cart">
  <h2>Shopping cart</h2>
  <ul>
    <li>
      <p><a href=""><strong>Silver earrings</strong></a>: $99.95.</p>
      <img src="../products/3333-0985/thumb.png" alt="Silver earrings">
    </li>
    <li>
      ...
    </li>
  </ul>
  <p>Total cost: $237.89</p>
</div>
```

This isn't really an ```<aside>```, as it doesn't necessarily relate to the main content of the page (you want it viewable from anywhere). It doesn't even particularly warrant using a  ```<section>```, as it isn't part of the main content of the page. So a ```<div>``` is fine in this case. We've included a heading as a signpost to aid screenreader users in finding it.

### Line breaks and horizontal rules

Two elements that you'll use occasionally and will want to know about are ```<br>``` and ```<hr>```.

__```<br>```__: __line break element__. ```<br>``` creates a line break in a paragraph; it is the only way to force a rigid structure in a situation where you want a series of fixed short lines, such as in a postal address or a poem. 
```html
<p>There once was a man named O'Dell<br>
Who loved to write HTML<br>
But his structure was bad, his semantics were sad<br>
and his markup didn't read very well.</p>
```

Without the ```<br>``` elements, the paragraph would just be rendered in one long line (as we said earlier in the course, HTML ignores most whitespace); 

__```<hr>```__: __thematic break element__. ```<hr>``` elements create a horizontal rule in the document that denotes a thematic change in the text (such as a change in topic or scene). Visually it just looks like a horizontal line.

```html
<p>Ron was backed into a corner by the marauding
   netherbeasts. Scared, but determined to protect his friends, he raised his
   wand and prepared to do battle, hoping that his distress call had made it through.</p>
<hr>
<p>Meanwhile, Harry was sitting at home, staring at his royalty statement
  and pondering when the next spin off series would come out, when an enchanted
  distress letter flew through his window and landed in his lap. He read it
  hazily and sighed; "better get back to work then", he mused.</p>
```

## Planning a Simple Website

See documentation for a good discussion: https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Document_and_website_structure#planning_a_simple_website

## Debugging HTML

HTML is not compiled into a different form before the browser parses it and shows the result (it is interpreted, not compiled). And HTML's element syntax is arguably a lot easier to understand than a "real programming language" like Rust, JavaScript, or Python. The way that browsers parse HTML is a lot more __permissive__ than how programming languages are run, which is both a good and a bad thing.

Generally when you do something wrong in code, there are two main types of error that you'll come across:

__Syntax errors__: These are spelling or punctuation errors in your code that actually cause the program not to run, like the Rust error shown above. These are usually easy to fix as long as you are familiar with the language's syntax and know what the error messages mean.

__Logic errors__: These are errors where the syntax is actually correct, but the code is not what you intended it to be, meaning that the program runs incorrectly. These are often harder to fix than syntax errors, as there isn't an error message to direct you to the source of the error.

__Permissive Code__: doesn't suffer from syntax errors

HTML itself doesn't suffer from syntax errors because browsers parse it <i>permissively</i>, meaning that the page still displays even if there are syntax errors. Browsers have built-in rules to state how to interpret incorrectly written markup, so you'll get something running, even if it is not what you expected. This, of course, can still be a problem!

### HTML Validation

So you can see from the above example that you really want to make sure your HTML is well-formed! But how? In a small example like the one seen above, it is easy to search through the lines and find the errors, but what about a huge, complex HTML document?

The best strategy is to start by running your HTML page through the <a href="https://validator.w3.org/">Markup Validation Service<a> — created and maintained by the W3C, the organization that looks after the specifications that define HTML, CSS, and other web technologies. This webpage takes an HTML document as an input, goes through it, and gives you a report to tell you what is wrong with your HTML.

## Multimedia and Embedding

See: https://developer.mozilla.org/en-US/docs/Learn/HTML/Multimedia_and_embedding

<a href="https://developer.mozilla.org/en-US/docs/Learn/HTML/Multimedia_and_embedding/Images_in_HTML">Images in HTML<a>

<a href="https://developer.mozilla.org/en-US/docs/Learn/HTML/Multimedia_and_embedding/Video_and_audio_content">Video and Audio in HTML<a>

<a href="https://developer.mozilla.org/en-US/docs/Learn/HTML/Multimedia_and_embedding/Other_embedding_technologies">From ```<object>``` to ```<iframe>```--other embedding technologies<a>

<a href="https://developer.mozilla.org/en-US/docs/Learn/HTML/Multimedia_and_embedding/Adding_vector_graphics_to_the_Web">Adding vector graphics to web<a>

<a href="https://developer.mozilla.org/en-US/docs/Learn/HTML/Multimedia_and_embedding/Responsive_images">Responsive Images<a>

## HTML tables

A very common task in HTML is structuring tabular data, and it has a number of elements and attributes for just this purpose. Coupled with a little CSS for styling, HTML makes it easy to display tables of information on the web such as your school lesson plan, the timetable at your local swimming pool, or statistics about your favorite dinosaurs or football team. This module takes you through all you need to know about structuring tabular data using HTML.

### HTML table basics

A table is a structured set of data made up of rows and columns (__tabular data__). A table allows you to quickly and easily look up values that indicate some kind of connection between different types of data, for example a person and their age, or a day of the week, or the timetable for a local swimming pool.

When implemented correctly, HTML tables are handled well by accessibility tools such as screen readers, so a successful HTML table should enhance the experience of sighted and visually impaired users alike.

<u>Styling</u>

<p align="center">
    <img src="imgs/table1.jpg" width="800">
</p>

You can also have a look at the live example on GitHub! One thing you'll notice is that the table does look a bit more readable there — this is because the table you see above on this page has minimal styling, whereas the GitHub version has more significant CSS applied.

<p align="center">
    <img src="imgs/table2.jpg" width="800">
</p>

Be under no illusion; for tables to be effective on the web, you need to provide some styling information with CSS, as well as good solid structure with HTML. In this module we are focusing on the HTML part; to find out about the CSS part you should visit our Styling tables article after you've finished here.

We won't focus on CSS in this module, but we have provided a minimal CSS stylesheet for you to use that will make your tables more readable than the default you get without any styling. You can find the stylesheet here, and you can also find an HTML template that applies the stylesheet — these together will give you a good starting point for experimenting with HTML tables.

```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Table template</title>
    <link href="minimal-table.css" rel="stylesheet" type="text/css">
  </head>
  <body>
    <h1>Table template</h1>


  </body>
</html>
```

```css
html {
  font-family: sans-serif;
}

table {
  border-collapse: collapse;
  border: 2px solid rgb(200,200,200);
  letter-spacing: 1px;
  font-size: 0.8rem;
}

td, th {
  border: 1px solid rgb(190,190,190);
  padding: 10px 20px;
}

th {
  background-color: rgb(235,235,235);
}

td {
  text-align: center;
}

tr:nth-child(even) td {
  background-color: rgb(250,250,250);
}

tr:nth-child(odd) td {
  background-color: rgb(245,245,245);
}

caption {
  padding: 10px;
}
```

__HTML tables should be used for tabular data__ — this is what they are designed for. Unfortunately, a lot of people used to use HTML tables to lay out web pages, e.g. one row to contain the header, one row to contain the content columns, one row to contain the footer, etc. You can find more details and an example at Page Layouts in our Accessibility Learning Module. This was commonly used because CSS support across browsers used to be terrible; table layouts are much less common nowadays, but you might still see them in some corners of the web.

<u>In short, using tables for layout rather than CSS layout techniques is a bad idea</u>. The main reasons are as follows:

1. __Layout tables reduce accessibility for visually impaired users__: Screenreaders, used by blind people, interpret the tags that exist in an HTML page and read out the contents to the user. Because tables are not the right tool for layout, and the markup is more complex than with CSS layout techniques, the screenreaders' output will be confusing to their users.
2. __Tables produce tag soup__: As mentioned above, table layouts generally involve more complex markup structures than proper layout techniques. This can result in the code being harder to write, maintain, and debug.
3. __Tables are not automatically responsive__: When you use proper layout containers (such as ```<header>```, ```<section>```, ```<article>```, or ```<div>```), their width defaults to 100% of their parent element. Tables on the other hand are sized according to their content by default, so extra measures are needed to get table layout styling to effectively work across a variety of devices.

### ```<table>```

__```<table>```__: __table element__

1. The content of every table is enclosed by these two tags : ```<table></table>```. Add these inside the body of your HTML.
2. The smallest container inside a table is a table cell, which is created by a ```<td>``` element ('td' stands for 'table data'). 

__```<td>```__: __table data__
```html
<!-- Making a row of 4 cells -->
<td>Hi, I'm your first cell.</td>
<td>I'm your second cell.</td>
<td>I'm your third cell.</td>
<td>I'm your fourth cell.</td>
```

As you will see, the cells are not placed underneath each other, rather they are automatically aligned with each other on the same row. Each ```<td>``` element creates a single cell and together they make up the first row. Every cell we add makes the row grow longer.

To stop this row from growing and start placing subsequent cells on a second row, we need to use the ```<tr>``` element ('tr' stands for 'table row'). 

__```<tr>```__: __table row__
```html
<!-- first row -->
<tr>
  <td>Hi, I'm your first cell.</td>
  <td>I'm your second cell.</td>
  <td>I'm your third cell.</td>
  <td>I'm your fourth cell.</td>
</tr>
<!-- second row -->
<tr>
  <td>Hi, I'm your first cell.</td>
  <td>I'm your second cell.</td>
  <td>I'm your third cell.</td>
  <td>I'm your fourth cell.</td>
</tr>
```

__Adding headers with ```<th>``` elements__

Now let's turn our attention to __table headers__ — special cells that go at the start of a row or column and define the type of data that row or column contains.

```html
<table>
  <tr>
    <td>&nbsp;</td>
    <td>Knocky</td>
    <td>Flor</td>
    <td>Ella</td>
    <td>Juan</td>
  </tr>
  <tr>
    <td>Breed</td>
    <td>Jack Russell</td>
    <td>Poodle</td>
    <td>Streetdog</td>
    <td>Cocker Spaniel</td>
  </tr>
  <tr>
    <td>Age</td>
    <td>16</td>
    <td>9</td>
    <td>10</td>
    <td>5</td>
  </tr>
  <tr>
    <td>Owner</td>
    <td>Mother-in-law</td>
    <td>Me</td>
    <td>Me</td>
    <td>Sister-in-law</td>
  </tr>
  <tr>
    <td>Eating Habits</td>
    <td>Eats everyone's leftovers</td>
    <td>Nibbles at food</td>
    <td>Hearty eater</td>
    <td>Will eat till he explodes</td>
  </tr>
</table>
```

<p align="center">
    <img src="imgs/table3.jpg" width="700">
</p>

__Allowing cells to span multiple rows and columns__

Sometimes we want cells to span multiple rows or columns. Take the following simple example, which shows the names of common animals. In some cases, we want to show the names of the males and females next to the animal name. Sometimes we don't, and in such cases we just want the animal name to span the whole table.

<p align="center">
    <img src="imgs/table4.jpg" width="200">
</p>

We need a way to get "Animals", "Hippopotamus", and "Crocodile" to span across two columns, and "Horse" and "Chicken" to span downwards over two rows. Fortunately, table headers and cells have the __```colspan```__ and __```rowspan```__ attributes, which allow us to do just those things. Both accept a unitless number value, which equals the number of rows or columns you want spanned. For example, colspan="2" makes a cell span two columns.

```html

<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Animals table</title>
    <link href="minimal-table.css" rel="stylesheet" type="text/css">
  </head>
  <body>
    <h1>Animals table</h1>

    <table>
      <tr>
        <th colspan="2">Animals</th>
      </tr>
      <tr>
        <th colspan="2">Hippopotamus</th>
      </tr>
      <tr>
        <th rowspan="2">Horse</th>
        <td>Mare</td>
      </tr>
      <tr>
        <td>Stallion</td>
      </tr>
      <tr>
        <th colspan="2">Crocodile</th>
      </tr>
      <tr>
        <th rowspan="2">Chicken</th>
        <td>Hen</td>
      </tr>
      <tr>
        <td>Rooster</td>
      </tr>
    </table>


  </body>
</html>
```

### Commong Styling to Columns (without ```col```)

HTML has a method of defining styling information for an entire column of data all in one place — the __```<col>```__ and __```<colgroup>```__ elements. These exist because it can be a bit annoying and inefficient having to specify styling on columns — you generally have to specify your styling information on every ```<td>``` or ```<th>``` in the column, or use a complex selector such as ```:nth-child```.

```html
<!-- not ideal. style every column by styling every cell -->
<table>
  <tr>
    <th>Data 1</th>
    <th style="background-color: yellow">Data 2</th>
  </tr>
  <tr>
    <td>Calcutta</td>
    <td style="background-color: yellow">Orange</td>
  </tr>
  <tr>
    <td>Robots</td>
    <td style="background-color: yellow">Jazz</td>
  </tr>
</table>
```

We have to repeat the styling information across all three cells in the column (__we'd probably have a class set on all three in a real project and specify the styling in a separate stylesheet__).

### Commong Styling to Columns (with ```col```)

Instead of doing this, we can specify the information once, on a __```<col>```__ element. ```<col>``` elements are  specified inside a __```<colgroup>```__ __container__ just below the opening ```<table>``` tag. We could create the same effect as we see above by specifying our table as follows

```html
<table>
    <!-- put the columns into groups and style each group -->
  <colgroup>
    <col>
    <col style="background-color: yellow">
  </colgroup>
  <tr>
    <th>Data 1</th>
    <th>Data 2</th>
  </tr>
  <tr>
    <td>Calcutta</td>
    <td>Orange</td>
  </tr>
  <tr>
    <td>Robots</td>
    <td>Jazz</td>
  </tr>
</table>

<!-- If we wanted to apply the styling information to both columns, we could just include one <col> element with a span attribute on it, like this: -->
<table>
    <!-- add span parameter -->
  <colgroup>
    <col style="background-color: yellow" span="2">
  </colgroup>
  <tr>
    <th>Data 1</th>
    <th>Data 2</th>
  </tr>
  <tr>
    <td>Calcutta</td>
    <td>Orange</td>
  </tr>
  <tr>
    <td>Robots</td>
    <td>Jazz</td>
  </tr>
</table>
```

### HTML table advanced features and accessibility

<u>Adding a caption to your table with ```<caption>```</u>

You can give your table a caption by putting it inside a __```<caption>```__ element and nesting that inside the ```<table>``` element. You should put it just below the opening ```<table>``` tag.

```html
<table>
  <caption>Dinosaurs in the Jurassic period</caption>

  ...
</table>
```

As you can infer from the brief example above, the caption is meant to contain a description of the table contents. This is useful for all readers wishing to get a quick idea of whether the table is useful to them as they scan the page, but particularly for blind users. Rather than have a screenreader read out the contents of many cells just to find out what the table is about, the user can rely on a caption and then decide whether or not to read the table in greater detail.

A caption is placed directly beneath the ```<table>``` tag.

<u>Adding structure with ```<thead>```, ```<tfoot>```, and ```<tbody>```</u>

As your tables get a bit more complex in structure, it is useful to give them more structural definition. One clear way to do this is by using __```<thead>```__, __```<tfoot>```__, and __```<tbody>```__, which allow you to mark up a __header__, __footer__, and __body__ section for the table.

These elements don't make the table any more accessible to screenreader users, and don't result in any visual enhancement on their own. They are however very useful for styling and layout — <u>acting as useful hooks for adding CSS to your table</u>. To give you some interesting examples, in the case of a long table you could make the table header and footer repeat on every printed page, and you could make the table body display on a single page and have the contents available by scrolling up and down.

To use them:

The ```<thead>``` element must wrap the part of the table that is the header — this is usually the first row containing the column headings, but this is not necessarily always the case. If you are using ```<col>```/```<colgroup>``` element, the table header should come just below those.
The ```<tfoot>``` element needs to wrap the part of the table that is the footer — this might be a final row with items in the previous rows summed, for example. You can include the table footer right at the bottom of the table as you'd expect, or just below the table header (the browser will still render it at the bottom of the table).
The ```<tbody>``` element needs to wrap the other parts of the table content that aren't in the table header or footer. It will appear below the table header or sometimes footer, depending on how you decided to structure it.
Note: ```<tbody>``` is always included in every table, implicitly if you don't specify it in your code. To check this, open up one of your previous examples that doesn't include ```<tbody>``` and look at the HTML code in your browser developer tools — you will see that the browser has added this tag for you. You might wonder why you ought to bother including it at all — you should, because it gives you more control over your table structure and styling.

Example
<p align="center">
    <img src="imgs/table5.jpg" width="700">
</p>

```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>My spending record</title>
    <link href="minimal-table.css" rel="stylesheet" type="text/css">
    <style>
    </style>
  </head>
  <body>
    <h1>My spending record</h1>

      <table>
        <caption>How I chose to spend my money</caption>
          <tr>
            <th>Purchase</th>
            <th>Location</th>
            <th>Date</th>
            <th>Evaluation</th>
            <th>Cost (€)</th>
          </tr>
          <tr>
            <td>SUM</td>
            <td>118</td>
          </tr>
          <tr>
            <td>Haircut</td>
            <td>Hairdresser</td>
            <td>12/09</td>
            <td>Great idea</td>
            <td>30</td>
          </tr>
          <tr>
            <td>Lasagna</td>
            <td>Restaurant</td>
            <td>12/09</td>
            <td>Regrets</td>
            <td>18</td>
          </tr>
          <tr>
            <td>Shoes</td>
            <td>Shoeshop</td>
            <td>13/09</td>
            <td>Big regrets</td>
            <td>65</td>
          </tr>
          <tr>
            <td>Toothpaste</td>
            <td>Supermarket</td>
            <td>13/09</td>
            <td>Good</td>
            <td>5</td>
          </tr>
    </table>

  </body>
</html>
```

```css
html {
  font-family: sans-serif;
}

table {
  border-collapse: collapse;
  border: 2px solid rgb(200,200,200);
  letter-spacing: 1px;
  font-size: 0.8rem;
}

td, th {
  border: 1px solid rgb(190,190,190);
  padding: 10px 20px;
}

th {
  background-color: rgb(235,235,235);
}

td {
  text-align: center;
}

tr:nth-child(even) td {
  background-color: rgb(250,250,250);
}

tr:nth-child(odd) td {
  background-color: rgb(245,245,245);
}

caption {
  padding: 10px;
}
```

Can <u>nest</u> tables

```html
<table id="table1">
  <tr>
    <th>title1</th>
    <th>title2</th>
    <th>title3</th>
  </tr>
  <tr>
    <td id="nested">
      <table id="table2">
        <tr>
          <td>cell1</td>
          <td>cell2</td>
          <td>cell3</td>
        </tr>
      </table>
    </td>
    <td>cell2</td>
    <td>cell3</td>
  </tr>
  <tr>
    <td>cell4</td>
    <td>cell5</td>
    <td>cell6</td>
  </tr>
</table>
```

<u>The ```scope``` attribute</u>

https://developer.mozilla.org/en-US/docs/Learn/HTML/Tables/Advanced#the_scope_attribute

<u>The  ```id``` and ```headers``` attributes</u>

An alternative to using the ```scope``` attribute is to use ```id``` and ```headers``` attributes to create associations between headers and cells. The way they are used is as follows:

1. You add a unique ```id``` to each ```<th>``` element.
2. You add a ```headers``` attribute to each ```<td>``` element. Each ```headers``` attribute has to contain a list of the ids of all the ```<th>``` elements that act as a header for that cell, separated by spaces.

This gives your HTML table an explicit definition of the position of each cell in the table, defined by the header(s) for each column and row it is part of, kind of like a spreadsheet. For it to work well, the table really needs both column and row headers.

Returning to our spending costs example, the previous two snippets could be rewritten like this:

```html
<thead>
  <tr>
    <th id="purchase">Purchase</th>
    <th id="location">Location</th>
    <th id="date">Date</th>
    <th id="evaluation">Evaluation</th>
    <th id="cost">Cost (€)</th>
  </tr>
</thead>
<tbody>
<tr>
  <th id="haircut">Haircut</th>
  <td headers="location haircut">Hairdresser</td>
  <td headers="date haircut">12/09</td>
  <td headers="evaluation haircut">Great idea</td>
  <td headers="cost haircut">30</td>
</tr>

  ...

</tbody>
```

> __Note: This method creates very precise associations between headers and data cells but it uses a lot more markup and does not leave any room for errors.  The scope approach is usually enough for most tables.__

# Cascading Stylesheets (CSS)

__CSS__ — is the first technology you should start learning after HTML. While HTML is used to define the structure and semantics of your content, <u>CSS is used to style it and lay it out</u>. For example, you can use CSS to alter the font, color, size, and spacing of your content, split it into multiple columns, or add animations and other decorative features.

## CSS first steps

CSS (Cascading Style Sheets) is used to style and lay out web pages — for example, to alter the font, color, size, and spacing of your content, split it into multiple columns, or add animations and other decorative features. This module provides a gentle beginning to your path towards CSS mastery with the basics of how it works, what the syntax looks like, and how you can start using it to add styling to HTML.

### What is CSS?

In the Introduction to HTML module we covered what HTML is, and how it is used to mark up documents. These documents will be readable in a web browser. Headings will look larger than regular text, paragraphs break onto a new line and have space between them. Links are colored and underlined to distinguish them from the rest of the text. <u>What you are seeing is the browser's default styles — very basic styles that the browser applies to HTML to make sure it will be basically readable even if no explicit styling is specified by the author of the page</u>.

However, the web would be a boring place if all websites looked like that. Using CSS you can control exactly how HTML elements look in the browser, presenting your markup using whatever design you like.

### What is CSS for?

As we have mentioned before, CSS is a language for specifying how documents are presented to users — how they are styled, laid out, etc.

A __document__ is usually a text file structured using a markup language — HTML is the most common markup language, but you may also come across other markup languages such as __SVG__ or __XML__.

__Presenting__ a document to a user means converting it into a form usable by your audience. Browsers, like Firefox, Chrome, or Edge , are designed to present documents visually, for example, on a computer screen, projector or printer.

CSS can be used for very basic document text styling — for example changing the color and size of headings and links. It can be used to create layout — for example turning a single column of text into a layout with a main content area and a sidebar for related information. It can even be used for effects such as animation. 

### CSS Syntax

CSS comments are done with  __```/*```__ and __```*/```__
```css
/* my first commnet */
```

CSS is a rule-based language — you define rules specifying groups of styles that should be applied to particular elements or groups of elements on your web page. For example "I want the main heading on my page to be shown as large red text."

The following code shows a very simple CSS rule that would achieve the styling described above:

```css
h1 {
    color: red;
    font-size: 5em;
}

/*
selector {
    property: value;
    property: value;
    property: value;
}
*/
```

The rule opens with a __selector__. This selects the HTML element that we are going to style. In this case we are styling level one headings (```<h1>```).

We then have a set of curly braces __```{ }```__. Inside those will be one or more __declarations__, which take the form of __property__ and __value__ pairs. Each pair specifies a property of the element(s) we are selecting, then a value that we'd like to give the property.

Before the colon, we have the property, and after the colon, the value. CSS properties have different allowable values, depending on which property is being specified. In our example, we have the color property, which can take various color values. We also have the font-size property. This property can take various size units as a value.

A CSS stylesheet will contain many such rules, written one after the other.

```css
h1 {
    color: red;
    font-size: 5em;
}

p {
    color: black;
}
```

You will find that you quickly learn some values, whereas others you will need to look up. <u>The individual property pages on MDN give you a quick way to look up properties and their values when you forget, or want to know what else you can use as a value</u>.

> Note: You can find links to all the CSS property pages (along with other CSS features) listed on the MDN <a href="https://developer.mozilla.org/en-US/docs/Web/CSS/Reference">CSS reference<a>. <u>Alternatively, you should get used to searching for "mdn css-feature-name" in your favorite search engine</u> whenever you need to find out more information about a CSS feature. For example, try searching for __"mdn color"__ and __"mdn font-size"__!

### CSS Modules

As there are so many things that you could style using CSS, the language is broken down into __modules__. You'll see reference to these modules as you explore MDN and many of the documentation pages are organized around a particular module. 

> For example, you could take a look at the MDN reference to the Backgrounds and Borders module to find out what its purpose is, and what different properties and other features it contains. You will also find links to the CSS Specification that defines the technology (see below).

At this stage you don't need to worry too much about how CSS is structured, however it can make it easier to find information if, for example, you are aware that a certain property is likely to be found among other similar things and are therefore probably in the same specification.

### CSS Specifications

All web standards technologies (HTML, CSS, JavaScript, etc.) are defined in giant documents called specifications (or "specs"), which are published by standards organizations (such as the W3C, WHATWG, ECMA, or Khronos) and define precisely how those technologies are supposed to behave.

CSS is no different — it is developed by a group within the W3C called the __CSS Working Group__. This group is made of representatives of browser vendors and other companies who have an interest in CSS. There are also other people, known as invited experts, who act as independent voices; they are not linked to a member organization.

__New CSS features are developed, or specified, by the CSS Working Group.__ Sometimes because a particular browser is interested in having some capability, other times because web designers and developers are asking for a feature, and sometimes because the Working Group itself has identified a requirement. CSS is constantly developing, with new features becoming available. However, a key thing about CSS is that everyone works very hard to never change things in a way that would break old websites. A website built in 2000, using the limited CSS available then, should still be usable in a browser today!

As a newcomer to CSS, it is likely that you will find the CSS specs overwhelming — they are intended for engineers to use to implement support for the features in user agents, not for web developers to read to understand CSS. Many experienced developers would much rather refer to MDN documentation or other tutorials. It is however worth knowing that they exist, understanding the relationship between the CSS you are using, browser support (see below), and the specs.

### Browser support information

Once CSS has been specified then it is only useful for us in developing web pages if one or more browsers have implemented it. This means that the code has been written to turn the instruction in our CSS file into something that can be output to the screen. We'll look at this process more in the lesson How CSS works. <u>It is unusual for all browsers to implement a feature at the same time, and so there is usually a gap where you can use some part of CSS in some browsers and not in others. For this reason, being able to check implementation status is useful.</u>

The __browser support status__ is shown on every MDN property page in a section named "Browser compatibility" (use this to check if the property can be used on your website). For example, the compatibility section for the CSS font-family property is reproduced below

<p align="center">
    <img src="imgs/support.jpg" width="700">
</p>

### Getting Started

Let's start with an html file
```html
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>Getting started with CSS</title>
</head>

<body>

    <h1>I am a level one heading</h1>

    <p>This is a paragraph of text. In the text is a <span>span element</span>
and also a <a href="https://example.com">link</a>.</p>

    <p>This is the second paragraph. It contains an <em>emphasized</em> element.</p>

    <ul>
        <li>Item <span>one</span></li>
        <li>Item two</li>
        <li>Item <em>three</em></li>
    </ul>

</body>

</html>
```

### Adding CSS

The very first thing we need to do is to tell the HTML document that we have some CSS rules we want it to use. __There are three different ways to apply CSS to an HTML document that you'll commonly come across, however, for now, we will look at the most usual and useful way of doing so — linking CSS from the head of your document.__

Create a file in the same folder as your HTML document and save it as styles.css. The .css extension shows that this is a CSS file.

To link styles.css to index.html add the following line somewhere inside the ```<head>``` of the HTML document

```html
<link rel="stylesheet" href="styles.css">
```

This ```<link>``` element tells the browser that we have a stylesheet, using the ```rel``` attribute, and the location of that stylesheet as the value of the ```href``` attribute. You can test that the CSS works by adding a rule to ```styles.css```. Using your code editor add the following to your CSS file:

```css
h1 {
  color: red;
}
```

Save your HTML and CSS files and reload the page in a web browser. The level one heading at the top of the document should now be red. If that happens, congratulations — you have successfully applied some CSS to an HTML document. If that doesn't happen, carefully check that you've typed everything correctly.

You can continue to work in ```styles.css``` locally, or you can use our interactive editor below to continue with this tutorial. The interactive editor acts as if the CSS in the first panel is linked to the HTML document, just as we have with our document above.

### Styling HTML elements

By making our heading red we have already demonstrated that we can target and style an HTML element. We do this by targeting an __element selector__ — <u>this is a selector that directly matches an HTML element name</u>. To target all paragraphs in the document you would use the selector ```p```. To turn all paragraphs green you would use:
```css
p {
  color: green;
}
```

__You can target multiple selectors at once__, by separating the selectors with a comma. If I want all paragraphs and all list items to be green my rule looks like this:

```css
p, li {
    color: green;
}
```

Try this out in the interactive editor below (edit the code boxes), or in your local CSS document.

### Changing the default behavior of elements

When we look at a well-marked up HTML document, even something as simple as our example, we can see how the browser is making the HTML readable by adding some default styling. Headings are large and bold and our list has bullets. This happens because browsers have internal stylesheets containing default styles, which they apply to all pages by default; without them all of the text would run together in a clump and we would have to style everything from scratch. All modern browsers display HTML content by default in pretty much the same way.

However, you will often want something other than the choice the browser has made. This can be done by choosing the HTML element that you want to change, and using a CSS rule to change the way it looks. A good example is our <ul>, an unordered list. It has list bullets, and if I decide I don't want those bullets I can remove them like so:

```css
li {
  list-style-type: none;
}
```

Try adding this to your CSS now.

The ```list-style-type``` property is a good property to look at on MDN to see which values are supported. Take a look at the page for <a href="https://developer.mozilla.org/en-US/docs/Web/CSS/list-style-type">list-style-type<a> and you will find an interactive example at the top of the page to try some different values in, then all allowable values are detailed further down the page.

Looking at that page you will discover that in addition to removing the list bullets you can change them — try changing them to square bullets by using a value of ```square```.

### Adding a class

So far we have styled elements based on their HTML element names. This works as long as you want all of the elements of that type in your document to look the same. Most of the time that isn't the case and so you will need to find a way to select a subset of the elements without changing the others. <u>The most common way to do this is to add a class to your HTML element and target that class</u>.

In your HTML document, add a ```class``` attribute to the second list item. Your list will now look like this:

```html
<ul>
  <li>Item one</li>
  <li class="special">Item two</li>
  <li>Item <em>three</em></li>
</ul>
```

In your CSS you can target the ```class``` of ```special``` by creating a selector that starts with a full stop character. Add the following to your CSS file:

```css
.special {
  color: orange;
  font-weight: bold;
}
```

Save and refresh to see what the result is.

You can apply the ```class``` of special to any element on your page that you want to have the same look as this list item. For example, you might want the ```<span>``` in the paragraph to also be orange and bold. Try adding a ```class``` of special to it, then reload your page and see what happens.

Sometimes you will see rules with a selector that lists the HTML element selector along with the class:

```css
li.special {
  color: orange;
  font-weight: bold;
}
```

This syntax means "target any li element that has a class of special". If you were to do this then you would no longer be able to apply the class to a ```<span>``` or another element by adding the class to it; you would have to add that element to the list of selectors:

```css
li.special,
span.special {
  color: orange;
  font-weight: bold;
}
```

As you can imagine, some classes might be applied to many elements and you don't want to have to keep editing your CSS every time something new needs to take on that style. <u>Therefore it is sometimes best to bypass the element and refer to the class, unless you know that you want to create some special rules for one element alone, and perhaps want to make sure they are not applied to other things.</u>


### Styling things based on their location in a document

There are times when you will want something to look different based on where it is in the document. There are a number of selectors that can help you here, but for now we will look at just a couple. In our document, there are two ```<em>``` elements — one inside a paragraph and the other inside a list item. To select only an ```<em>``` that is nested inside an ```<li>``` element I can use a selector called the __descendant combinator__, which takes the form of a space between two other selectors.

Add the following rule to your stylesheet.
```css
/* descendent combinator: apply to any em elements inside an li element */
li em {
  color: rebeccapurple;
}
```

This selector will select any ```<em>``` element that is inside (a descendant of) an ```<li>```. So in your example document, you should find that the ```<em>``` in the third list item is now purple, but the one inside the paragraph is unchanged.

Something else you might like to try is styling a paragraph when it comes directly after a heading at the same hierarchy level in the HTML. To do so place a __```+```__ (an __adjacent sibling combinator__) between the selectors.

Try adding this rule to your stylesheet as well:

```css
/* adjacent sibling combinator: apply to any paragraph after a header1*/
h1 + p {
  font-size: 200%;
}
```

### Styling things based on state

The final type of styling we shall take a look at in this tutorial is the ability to style things based on their __state__. A straightforward example of this is when styling links. When we style a link we need to target the ```<a>``` (anchor) element. This has different states depending on whether it is <u>unvisited</u>, <u>visited</u>, being </u>hovered over</u>, <u>focused via the keyboard</u>, or in the process of being clicked (activated). You can use CSS to target these different states — the CSS below styles unvisited links pink and visited links green.

```css
/* state: apply to the following elements based on their state */
a:link {
  color: pink;
}

a:visited {
  color: green;
}
```

You can change the way the link looks when the user hovers over it, for example by removing the underline, which is achieved by the next rule:

```css
a:hover {
  text-decoration: none;
}
```

### Combining selectors and combinators

It is worth noting that you can combine multiple selectors and combinators together. 

```css
/* selects any <span> that is inside a <p>, which is inside an <article>  */
article p span { ... }

/* selects any <p> that comes directly after a <ul>, which comes directly after an <h1>  */
h1 + ul + p { ... }
```

You can combine multiple types together, too. Try adding the following into your code:

```css
body h1 + p .special {
  color: yellow;
  background-color: black;
  padding: 5px;
}
```

This will style any element with a class of special, which is inside a ```<p>```, which comes just after an ```<h1>```, which is inside a ```<body>```. Phew!

In the original HTML we provided, the only element styled is ```<span class="special">```.

Don't worry if this seems complicated at the moment — you'll soon start to get the hang of it as you write more CSS.

### How CSS is structured

### Applying CSS to HTML

First, let's examine three methods of applying CSS to a document:

1. __with an external stylesheet__
2. __with an internal stylesheet__
3. __with inline styles__

### CSS External stylesheet

An external stylesheet contains CSS in a separate file with a ```.css``` extension. This is the most common and useful method of bringing CSS to a document. You can link a single CSS file to multiple web pages, styling all of them with the same CSS stylesheet. In the Getting started with CSS, we linked an external stylesheet to our web page.

You reference an external CSS stylesheet from an HTML ```<link>``` element:

```html
<!-- Example of external CSS styling-->
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>My CSS experiment</title>
    <!-- add link to css style sheet here! -->
    <link rel="stylesheet" href="styles.css">
  </head>
  <body>
    <h1>Hello World!</h1>
    <p>This is my first CSS example</p>
  </body>
</html>
```

The CSS stylesheet file might look like this:

```css
/* Example of external CSS styling */
h1 {
  color: blue;
  background-color: yellow;
  border: 1px solid black;
}

p {
  color: red;
}
```

The href attribute of the __```<link>```__ element needs to reference a file on your file system. In the example above, the CSS file is in the same folder as the HTML document, but you could place it somewhere else and adjust the path. Here are three examples:

```html
<!-- Inside a subdirectory called styles inside the current directory -->
<link rel="stylesheet" href="styles/style.css">

<!-- Inside a subdirectory called general, which is in a subdirectory called styles, inside the current directory -->
<link rel="stylesheet" href="styles/general/style.css">

<!-- Go up one directory level, then inside a subdirectory called styles -->
<link rel="stylesheet" href="../styles/style.css">
```

### CSS Internal Stylesheet

An internal stylesheet resides within an HTML document. To create an internal stylesheet, you place CSS inside a __```<style>```__ element contained inside the HTML ```<head>```.

The HTML for an internal stylesheet might look like this:

```html
<!-- Example of internal CSS stylesheet-->
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>My CSS experiment</title>
    <style>
      h1 {
        color: blue;
        background-color: yellow;
        border: 1px solid black;
      }

      p {
        color: red;
      }
    </style>
  </head>
  <body>
    <h1>Hello World!</h1>
    <p>This is my first CSS example</p>
  </body>
</html>
```

__In some circumstances, internal stylesheets can be useful__. For example, perhaps you're working with a content management system where you are blocked from modifying external CSS files.

__But for sites with more than one page, an internal stylesheet becomes a less efficient way of working__. To apply uniform CSS styling to multiple pages using internal stylesheets, you must have an internal stylesheet in every web page that will use the styling. __The efficiency penalty carries over to site maintenance too. With CSS in internal stylesheets, there is the risk that even one simple styling change may require edits to multiple web pages.__


### CSS Inline Styles

Inline styles are CSS declarations that affect a single HTML element, contained within a style attribute. The implementation of an inline style in an HTML document might look like this:

```html
<!-- Example of inline CSS styling-->
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>My CSS experiment</title>
  </head>
  <body>
      <!-- inline style here -->
    <h1 style="color: blue;background-color: yellow;border: 1px solid black;">Hello World!</h1>
    <p style="color:red;">This is my first CSS example</p>
  </body>
</html>
```

__Avoid using CSS in this way, when possible__. It is the opposite of a best practice. First, it is the least efficient implementation of CSS for maintenance. One styling change might require multiple edits within a single web page. Second, inline CSS also mixes (CSS) presentational code with HTML and content, making everything more difficult to read and understand. Separating code and content makes maintenance easier for all who work on the website.

There are a few circumstances where inline styles are more common. __You might have to resort to using inline styles if your working environment is very restrictive__. For example, perhaps your CMS only allows you to edit the HTML body. __You may also see a lot of inline styles in HTML email to achieve compatibility with as many email clients as possible__.

### CSS Selectors

A selector targets HTML to apply styles to content. We have already discovered a variety of selectors in the Getting started with CSS section. <u>If CSS is not applying to content as expected, your selector may not match the way you think it should match</u>.

<u>Each CSS rule starts with a selector — or a list of selectors — in order to tell the browser which element or elements the rules should apply to</u>. All the examples below are valid selectors or lists of selectors.

```css
h1
a:link
.manythings
#onething
*
.box p
.box p:first-child
h1, h2, .intro
```

### CSS Selector (Specificity)

You may encounter scenarios where two selectors select the same HTML element. Consider the stylesheet below, with a ```p``` selector that sets paragraph text to blue. However, there is also a ```class``` that sets the text of selected elements to red.

```css
.special {
  color: red;
}

p {
  color: blue;
}
```

Suppose that in our HTML document, we have a paragraph with a class of special. Both rules apply. Which selector prevails? Do you expect to see blue or red paragraph text?

```html
<p class="special">What color am I?</p>
```

<u>The CSS language has rules to control which selector is stronger in the event of a conflict</u>. These rules are called __cascade__ and __specificity__. In the code block below, we define two rules for the p selector, but the paragraph text will be blue. This is because the declaration that sets the paragraph text to blue appears later in the stylesheet. Later styles replace conflicting styles that appear earlier in the stylesheet. This is the cascade rule.

__Cascade rule__: <u>Later styles replace conflicting styles that appear earlier in the stylesheet</u>

```css
p {
  color: red;
}

p {
  color: blue;
}
```

__However, in the case of our earlier example with the conflict between the class selector and the element selector, the class prevails, rendering the paragraph text red__. How can this happen even though a conflicting style appears later in the stylesheet? __A class is rated as being more specific, as in having more specificity than the element selector, so it cancels the other conflicting style declaration.__

__Specificity rule__: <u>selectors that are more specific (e.g. class selector being more specific than an element selector) take precidence</u>

The rules of specificity and the cascade can seem complicated at first. These rules are easier to understand as you become more familiar with CSS. The Cascade and inheritance section in the next module explains this in detail, including how to calculate specificity.

For now, remember that specificity exists. Sometimes, CSS might not apply as you expected because something else in the stylesheet has more specificity. Recognizing that more than one rule could apply to an element is the first step in fixing these kinds of issues.

### CSS Properties and values

At its most basic level, CSS consists of two components:

1. __Properties__: These are human-readable identifiers that indicate which stylistic features you want to modify. 
>For example, font-size, width, background-color.

2. __Values__: Each property is assigned a value. This value indicates how to style the property.

The example below highlights a single property and value. The property name is color and the value is blue.

```css
h1 {
    color: blue;
    background-color: yellow;
}
p {
    /* property: value;*/
    color: blue;
}
```

Finally, __CSS declaration blocks__ are paired with __selectors__ to produce __CSS rulesets__ (or __CSS rules__). The example below contains two rules: one for the h1 selector and one for the p selector. The colored highlighting identifies the h1 rule.

```css
/* CSS rule 1*/

/* CSS selector(s) */ 
h1 {
    /* CSS declaration block */
    color: blue;
    background-color: yellow;
}
/* CSS rule 2*/

/* CSS selector(s) */ 
p {
    /* CSS declaration block */
    color: blue;
}
```

Setting CSS properties to specific values is the primary way of defining layout and styling for a document. The CSS engine calculates which declarations apply to every single element of a page.

__Warning: CSS properties and values are case-sensitive. The property and value in each pair are separated by a colon. (```:```)__

Look up different values of properties listed below. Write CSS rules that apply styling to different HTML elements:

<a href="https://developer.mozilla.org/en-US/docs/Web/CSS/font-size">font-size<a>

<a href="https://developer.mozilla.org/en-US/docs/Web/CSS/width">width<a>

<a href="https://developer.mozilla.org/en-US/docs/Web/CSS/background-color">background-color<a>

<a href="https://developer.mozilla.org/en-US/docs/Web/CSS/color">color<a>

<a href="https://developer.mozilla.org/en-US/docs/Web/CSS/border">border<a>

### CSS Functions
While most values are relatively simple keywords or numeric values, there are some values that take the form of a function.

__The ```calc()``` function__
An example would be the calc() function, which can do simple math within CSS:

```html
<div class="outer"><div class="box">The inner box is 90% - 30px.</div></div>
```

```css
.outer {
  border: 5px solid black;
}

.box {
  padding: 10px;
  width: calc(90% - 30px);
  background-color: rebeccapurple;
  color: white;
}
```

<p align="center">
    <img src="imgs/func1.jpg" width="600">
</p>


A function consists of the function name, and parentheses to enclose the values for the function. In the case of the ```calc()``` example above, the values define the width of this box to be 90% of the containing block width, minus 30 pixels. The result of the calculation isn't something that can be computed in advance and entered as a static value.

### CSS Transform functions
Another example would be the various values for transform, such as ```rotate()```.

```html
<div class="box"></div>
```
```css
.box {
  margin: 30px;
  width: 100px;
  height: 100px;
  background-color: rebeccapurple;
  transform: rotate(0.8turn);
}
```
<p align="center">
    <img src="imgs/func2.JPG" width="200">
</p>

Look up different values of properties listed below. Write CSS rules that apply styling to different HTML elements:

<a href="https://developer.mozilla.org/en-US/docs/Web/CSS/transform">transform <a>

<a href="https://developer.mozilla.org/en-US/docs/Web/CSS/background-image">background-image<a>, in particular gradient values

<a href="https://developer.mozilla.org/en-US/docs/Web/CSS/color">color<a>, in particular rgb/rgba/hsl/hsla values

### CSS @rules

__CSS @rules__ (pronounced "at-rules") provide instruction for what CSS should perform or how it should behave. Some @rules are simple with just a keyword and a value. For example, __```@import```__ imports a stylesheet into another CSS stylesheet:

```css
@import 'styles2.css';
```

One common @rule that you are likely to encounter is __```@media```__, which is used to create media queries. Media queries use conditional logic for applying CSS styling.

In the example below, the stylesheet defines a default pink background for the ```<body>``` element. However, a media query follows that defines a blue background if the browser viewport is wider than ```30em```.

```css
body {
  background-color: pink;
}

@media (min-width: 30em) {
  body {
    background-color: blue;
  }
}
```

### CSS Shorthands

Some properties like ```font```, ```background```, ```padding```, ```border```, and ```margin``` are called __shorthand properties__. <u>This is because shorthand properties set several values in a single line</u>.

For example, this one line of code:

```css
/* In 4-value shorthands like padding and margin, the values are applied
   in the order top, right, bottom, left (clockwise from the top). There are also other
   shorthand types, for example 2-value shorthands, which set padding/margin
   for top/bottom, then left/right */
padding: 10px 15px 15px 5px;
```

is equivalent to these four lines of code:

```css
padding-top: 10px;
padding-right: 15px;
padding-bottom: 15px;
padding-left: 5px;
```

This one line:

```css
background: red url(bg-graphic.png) 10px 10px repeat-x fixed;
```

is equivalent to these five lines:

```css
background-color: red;
background-image: url(bg-graphic.png);
background-position: 10px 10px;
background-repeat: repeat-x;
background-attachment: fixed;
```

Later in the course, you will encounter many other examples of shorthand properties. MDN's <a href="https://developer.mozilla.org/en-US/docs/Web/CSS/Reference">CSS reference<a> is a good resource for more information about any shorthand property.

Try using the declarations (above) in your own CSS exercise to become more familiar with how it works. You can also experiment with different values.

### CSS Comments

As with any coding work, it is best practice to write __comments__ along with CSS. This helps you to remember how the code works as you come back later for fixes or enhancement. It also helps others understand the code.

CSS comments begin with __```/*```__ and end with __```*/```__. In the example below, comments mark the start of distinct sections of code. This helps to navigate the codebase as it gets larger. With this kind of commenting in place, searching for comments in your code editor becomes a way to efficiently find a section of code.

```css
/* Handle basic element styling */
/* -------------------------------------------------------------------------------------------- */
body {
  font: 1em/150% Helvetica, Arial, sans-serif;
  padding: 1em;
  margin: 0 auto;
  max-width: 33em;
}

@media (min-width: 70em) {
  /* Increase the global font size on larger screens or windows
     for better readability */
  body {
    font-size: 130%;
  }
}

h1 {font-size: 1.5em;}

/* Handle specific elements nested in the DOM  */
/* -------------------------------------------------------------------------------------------- */
div p, #id:first-line {
  background-color: red;
  border-radius: 3px;
}

div p {
  margin: 0;
  padding: 1em;
}

div p + p {
  padding-top: 0;
}
```

"Commenting out" code is also useful for temporarily disabling sections of code for testing. In the example below, the rules for .special are disabled by "commenting out" the code.

```css
/*.special {
  color: red;
}*/

p {
  color: blue;
}
```

### CSS White space

White space means actual spaces, tabs and new lines. Just as browsers ignore white space in HTML, <u>browsers ignore white space inside CSS</u>. The value of white space is how it can improve readability.

In the example below, each declaration (and rule start/end) has its own line. This is arguably a good way to write CSS. It makes it easier to maintain and understand CSS.

```css
body {
  font: 1em/150% Helvetica, Arial, sans-serif;
  padding: 1em;
  margin: 0 auto;
  max-width: 33em;
}

@media (min-width: 70em) {
  body {
    font-size: 130%;
  }
}

h1 {
  font-size: 1.5em;
}

div p,
#id:first-line {
  background-color: red;
  border-radius: 3px;
}

div p {
  margin: 0;
  padding: 1em;
}

div p + p {
  padding-top: 0;
}
```

> __Warning: Though white space separates values in CSS declarations, property names never have white space.__

For example, these declarations are valid CSS:

```css
margin: 0 auto;
padding-left: 10px;
```

But these declarations are invalid:

```css
margin: 0auto;
padding- left: 10px;

/* error because property has a space in it!*/
```

### How does CSS actually work?

When a browser displays a document, it must combine the document's content with its style information. It processes the document in a number of stages, which we've listed below. Bear in mind that this is a very simplified version of what happens when a browser loads a webpage, and that different browsers will handle the process in different ways. But this is roughly what happens.

1. The browser loads the HTML (e.g. receives it from the network).
2. It converts the HTML into a __DOM (Document Object Model)__. The DOM represents the document in the computer's memory. The DOM is explained in a bit more detail in the next section.
3. The browser then fetches most of the resources that are linked to by the HTML document, such as embedded images and videos ... and linked CSS! JavaScript is handled a bit later on in the process, and we won't talk about it here to keep things simpler.
4. The browser parses the fetched CSS, and sorts the different rules by their selector types into different "buckets", e.g. element, class, ID, and so on. Based on the selectors it finds, it works out which rules should be applied to which nodes in the DOM, and attaches style to them as required (this intermediate step is called a render tree).
5. The render tree is laid out in the structure it should appear in after the rules have been applied to it.
6. The visual display of the page is shown on the screen (this stage is called painting).

The following diagram also offers a simple view of the process.

<p align="center">
    <img src="imgs/rendering.svg" width="800">
</p>

For more on DOM: https://developer.mozilla.org/en-US/docs/Learn/CSS/First_steps/How_CSS_works#about_the_dom

### HTML, CSS Example



<p align="center">
    <img src="imgs/resume_ex.jpg" width="800">
</p>

```css
body {
    font-family: Arial, Helvetica, sans-serif;
}

h1 {
    color: #375e97;
    font-size: 2em;
    font-family: Georgia, 'Times New Roman', Times, serif;
    border-bottom: 1px solid #375e97;
}

h2 {
    font-size: 1.5em;
}

.job-title {
    color: #999999;
    font-weight: bold;
}

a:link, a:visited {
    color: #fb6542;
}

a:hover {
    text-decoration: none;
} 
```

```html
<h1>Jane Doe</h1>
<div class="job-title">Web Developer</div>
<p>Far far away, behind the word mountains, far from the countries Vokalia and Consonantia, there live the blind texts. Separated they live in Bookmarksgrove right at the coast of the Semantics, a large language ocean.</p>

<p>A small river named Duden flows by their place and supplies it with the necessary regelialia. It is a paradisematic country, in which roasted parts of sentences fly into your mouth. </p>

<h2>Contact information</h2>
<ul>
    <li>Email: <a href="mailto:jane@example.com">jane@example.com</a></li>
    <li>Web: <a href="http://example.com">http://example.com</a></li>
    <li>Tel: 123 45678</li>
</ul>
    z
```


## CSS building blocks

### Conflicting Rules: Cascading, Specificity

__Cascade__

Stylesheets cascade — at a very simple level, this means that the order of CSS rules matters; when two rules apply that have equal specificity, the one that comes last in the CSS is the one that will be used.

In the below example, we have two rules that could apply to the h1. The h1 ends up being colored blue — these rules have an identical selector and therefore carry the same specificity, so the last one in the source order wins.

__Specificity__

Specificity is how the browser decides which rule applies if multiple rules have different selectors, but could still apply to the same element. It is basically a measure of how specific a selector's selection will be:

- __An element selector is less specific__ — it will select all elements of that type that appear on a page — so will get a lower score.
- __A class selector is more specific__ — it will select only the elements on a page that have a specific class attribute value — so will get a higher score.

__Inheritance__

Inheritance also needs to be understood in this context — some CSS property values set on parent elements are inherited by their child elements, and some aren't.

For example, if you set a ```color``` and ```font-family``` on an element, every element inside it will also be styled with that color and font, unless you've applied different color and font values directly to them.

```css
/* */
body {
    color: blue;
}

span {
    color: black;
}
```

__Understanding inheritance__

We'll start with inheritance. In the example below we have a ```<ul>```, with two levels of unordered lists nested inside it. We have given the outer ```<ul>``` a border, padding, and font color.

The color has applied to the direct children, but also the indirect children — the immediate child ```<li>```s, and those inside the first nested list. We have then added a class- special to the second nested list and applied a different color to it. This then inherits down through its children.

```css
.main {
    color: rebeccapurple;
    border: 2px solid #ccc;
    padding: 1em;
}

.special {
    color: black;
    font-weight: bold;
}
```

```html
<ul class="main">
    <li>Item One</li>
    <li>Item Two
        <ul>
            <li>2.1</li>
            <li>2.2</li>
        </ul>
    </li>
    <li>Item Three
        <ul class="special">
            <li>3.1
                <ul>
                    <li>3.1.1</li>
                    <li>3.1.2</li>
                </ul>
            </li>
            <li>3.2</li>
        </ul>
    </li>
</ul>
```

Things like widths (as mentioned above), margins, padding, and borders do not inherit. If a border were to be inherited by the children of our list, every single list and list item would gain a border — probably not an effect we would ever want!

Which properties are inherited by default and which aren't is largely down to common sense.


__Controlling inheritance__

CSS provides four special universal property values for controlling inheritance. Every CSS property accepts these values.

__```inherit```__: Sets the property value applied to a selected element to be the same as that of its parent element. Effectively, this "turns on inheritance".

__```initial```__: Sets the property value applied to a selected element to the initial value of that property.

__```unset```__: Resets the property to its natural value, which means that if the property is naturally inherited it acts like ```inherit```, otherwise it acts like ```initial```.

__```revert```__: Acts like ```unset``` in many cases, however will revert the property to the browser's default styling rather than the defaults applied to that property.

__Resetting all property values__

The CSS shorthand property __```all```__ can be used to apply one of these inheritance values to (almost) all properties at once. Its value can be any one of the inheritance values (```inherit```, ```initial```, ```unset``` or ```revert```). It's a convenient way to undo changes made to styles so that you can get back to a known starting point before beginning new changes.

In the below example we have two blockquotes. The first has styling applied to the blockquote element itself, the second has a class applied to the blockquote which sets the value of all to unset.

```css
blockquote {
    background-color: red;
    border: 2px solid green;
}
        
.fix-this {
    all: unset;
}
```

```html
        <blockquote>
            <p>This blockquote is styled</p>
        </blockquote>

        <blockquote class="fix-this">
            <p>This blockquote is not styled</p>
        </blockquote>
```

__Understanding the cascade__

We now understand why a paragraph nested deep in the structure of your HTML is the same color as the CSS applied to the body, and from the introductory lessons, we have an understanding of how to change the CSS applied to something at any point in the document — whether by assigning CSS to an element or creating a class. We will now take a proper look at how the cascade defines which CSS rules apply when more than one thing could style an element.

There are three factors to consider, listed here in increasing order of importance. Later ones overrule earlier ones:

1. __Source order__
2. __Specificity__
3. __Importance__
We will look at these to see how browsers figure out exactly what CSS should be applied.

__Source order__: We have already seen how source order matters to the cascade. If you have more than one rule, which has exactly the same weight, then the one that comes last in the CSS will win. You can think of this as rules which are nearer the element itself overwriting early ones until the last one wins and gets to style the element.

__Specificity__: Once you understand the fact that source order matters, at some point you will run into a situation where you know that a rule comes later in the stylesheet, but an earlier, conflicting, rule is applied. This is because that earlier rule has a higher specificity — it is more specific, and therefore is being chosen by the browser as the one that should style the element.

> As we saw earlier in this lesson, a class selector has more weight than an element selector, so the properties defined on the class will override those applied directly to the element.

Something to note here is that although we are thinking about selectors and the rules that are applied to the thing they select, it isn't the entire rule which is overwritten, only the properties which are the same.

This behavior helps avoid repetition in your CSS. <u>A common practice is to define generic styles for the basic elements, and then create classes for those which are different</u>. 

For example, in the stylesheet below we have defined generic styles for level 2 headings, and then created some classes which change only some of the properties and values. The values defined initially are applied to all headings, then the more specific values are applied to the headings with the classes.

```css
h2 {
    font-size: 2em;
    color: #000;
    font-family: Georgia, 'Times New Roman', Times, serif;
}
        
.small {
    font-size: 1em;
}
        
.bright {
    color: rebeccapurple;
}         
```

```html
<h2>Heading with no class</h2>
<h2 class="small">Heading with class of small</h2>
<h2 class="bright">Heading with class of bright</h2>
```

__```!important```__

There is a special piece of CSS that you can use to overrule all of the above calculations. However, you should be very careful with using it — ```!important```. This is used to make a particular property and value the most specific thing, thus overriding the normal rules of the cascade.

Take a look at this example where we have two paragraphs, one of which has an ID.

```css
#winning {
    background-color: red;
    border: 1px solid black;
}
    
.better {
    background-color: gray;
    border: none !important;
}
    
p {
    background-color: blue;
    color: white;
    padding: 5px;
}     
```

```html
<p class="better">This is a paragraph.</p>
<p class="better" id="winning">One selector to rule them all!</p>
```

### CSS Selectors

__CSS Selector (lists)__

If you have more than one thing which uses the same CSS then the individual selectors can be combined into a selector list so that the rule is applied to all of the individual selectors. For example, if I have the same CSS for an h1 and also a class of .special, I could write this as two separate rules.

```css
h1 {
  color: blue;
}

.special {
  color: blue;
}
```

I could also combine these into a selector list, by adding a comma between them.

```css
h1, .special {
  color: blue;
}
```

White space is valid before or after the comma. You may also find the selectors more readable if each is on a new line.

```css
h1,
.special {
  color: blue;
}
```

### CSS Selectors: Type, class, and ID selectors

This group includes selectors that target an HTML element such as an ```<h1>```.

```css
h1 { }
```

It also includes selectors which target a class:

```css
.box { }
```

or, an ID:

```css
#unique { }
```

### CSS Selectors: Attribute selectors

This group of selectors gives you different ways to select elements based on the presence of a certain attribute on an element:

```css
/* select anchor elements with title attributes? */
a[title] { }
```

Or even make a selection based on the presence of an attribute with a particular value:

```css
/* select anchor elements that link to this place */
a[href="https://example.com"] { }
```

### CSS Selectors: Pseudo-classes and pseudo-elements

This group of selectors includes __pseudo-classes__, which __style certain states of an element__. The ```:hover``` pseudo-class for example selects an element only when it is being hovered over by the mouse pointer:

```css
a:hover { }
```

It also includes __pseudo-elements__, which __select a certain part of an element rather than the element itself__. For example, ```::first-line``` always selects the first line of text inside an element (a ```<p>``` in the below case), acting as if a ```<span>``` was wrapped around the first formatted line and then selected.

```css
p::first-line { }
```

### CSS Selectors: Combinators

The final group of selectors __combine other selectors in order to target elements within our documents__. The following, for example, selects paragraphs that are direct children of ```<article>``` elements using the child combinator (__```>```__):

```css
article > p { }
```

### CSS Selectors: table summary

<p align="center">
    <img src="imgs/selectors_and_combinators.jpg" width=100%>
</p>

### CSS: Box Model

In CSS we broadly have two types of boxes — __block boxes__ and __inline boxes__. These characteristics refer to how the box behaves in terms of page flow and in relation to other boxes on the page. Boxes also have an __inner display__ type and an __outer display__ type. First, we will explain what we mean by block box and inline box. We will then explain what is meant by an inner and outer display type.

If a box has an outer display type of ```block```, it will behave in the following ways:

- The box will break onto a new line.
- The box will extend in the inline direction to fill the space available in its container. In most cases this means that the box will become as wide as its container, filling up 100% of the space available.
- The ```width``` and ```height``` properties are respected.
- Padding, margin and border will cause other elements to be pushed away from the box

> Some HTML elements, such as ```<h1>``` and ```<p>```, use block as their outer display type by default.

If a box has an outer display type of ```inline```, then:

- The box will not break onto a new line.
- The ```width``` and ```height``` properties will not apply.
- Vertical padding, margins, and borders will apply but will not cause other inline boxes to move away from the box.
- Horizontal padding, margins, and borders will apply and will cause other inline boxes to move away from the box.

> Some HTML elements, such as ```<a>```, ```<span>```, ```<em>``` and ```<strong>``` use inline as their outer display type by default.

The type of box applied to an element is defined by <a href="https://developer.mozilla.org/en-US/docs/Web/CSS/display">```display```<a> property values such as ```block``` and ```inline```, and relates to the outer value of display.

### CSS: Box Model: Aside: Inner and outer display types

At this point, we'd better also explain __inner__ and outer __display__ types. As mentioned above, boxes in CSS have an outer display type, which details whether the box is block or inline.

Boxes also have an inner display type, however, which dictates how elements inside that box are laid out. By default, the elements inside a box are laid out in __normal flow__, which means that they behave just like any other block and inline elements (as explained above).

We can, however, change the inner display type by using ```display``` values like ```flex```. If we set ```display: flex;``` on an element, the outer display type is block, but the inner display type is changed to flex. Any direct children of this box will become flex items and will be laid out according to the rules set out in the Flexbox spec, which you'll learn about later on

> Note: To read more about the values of display, and how boxes work in block and inline layout, take a look at the MDN guide to Block and Inline Layout.

When you move on to learn about CSS Layout in more detail, you will encounter ```flex```, and various other inner values that your boxes can have, for example ```grid```.

Block and inline layout, however, is the default way that things on the web behave — as we said above, it is sometimes referred to as normal flow, because without any other instruction, our boxes lay out as block or inline boxes.

### CSS: Box Model: different display types

Below we have three different HTML elements, all of which have an outer display type of ```block```. The first is a paragraph, which has a border added in CSS. The browser renders this as a block box, so the paragraph begins on a new line, and expands to the full width available to it.

The second is a list, which is laid out using ```display: flex```. This establishes flex layout for the items inside the container, however, the list itself is a block box and — like the paragraph — expands to the full container width and breaks onto a new line.

Below this, we have a block-level paragraph, inside which are two ```<span>``` elements. These elements would normally be ```inline```, however, one of the elements has a class of block, and we have set it to ```display: block```.



<p align="center">
    <img src="imgs/display_block.jpg" width=100%>
</p>

```css
p, 
ul {
  border: 2px solid rebeccapurple;
  padding: .5em;
}

.block,
li {
  border: 2px solid blue;
  padding: .5em;
}

ul {
  display: flex;
  list-style: none;
}

.block {
  display: block;
}      
```   
```html
<p>I am a paragraph. A short one.</p>
<ul>
  <li>Item One</li>
  <li>Item Two</li>
  <li>Item Three</li>
</ul>
<p>I am another paragraph. Some of the <span class="block">words</span> have been wrapped in a <span>span element</span>.</p>
```

### CSS: Box Model

The CSS __box model__ as a whole applies to block boxes. Inline boxes use just some of the behavior defined in the box model. The model defines how the different parts of a box — margin, border, padding, and content — work together to create a box that you can see on a page. To add some additional complexity, there is a __standard box model__ and an __alternate box model__.

#### Parts of a box

Making up a block box in CSS we have the:

__Content box__: The area where your content is displayed, which can be sized using properties like ```width``` and ```height```.\
__Padding box__: The padding sits around the content as white space; its size can be controlled using ```padding``` and related properties.\
__Border box__: The border box wraps the content and any padding. Its size and style can be controlled using ```border``` and related properties.\
__Margin box__: The margin is the outermost layer, wrapping the content, padding, and border as whitespace between this box and other elements. Its size can be controlled using ```margin``` and related properties.

The below diagram shows these layers:

<p align="center">
    <img src="imgs/box-model.png" width=80%>
</p>


#### Standard CSS box Model

In the standard box model, if you give a box a ```width``` and a ```height``` attribute, this defines the width and height of the content box. Any padding and border is then added to that width and height to get the total size taken up by the box. This is shown in the image below.

If we assume that a box has the following CSS defining ```width```, ```height```, ```margin```, ```border```, and ```padding```:

```css
/* actual space taken up will be width + 2*padding +2*border */
/* margin not included */
.box {
  width: 350px;
  height: 150px;
  margin: 10px;
  padding: 25px;
  border: 5px solid black;
}
```

<p align="center">
    <img src="imgs/standard-box-model.png" width=80%>
</p>


> __Note__: The margin is not counted towards the actual size of the box — sure, it affects the total space that the box will take up on the page, but only the space outside the box. The box's area stops at the border — it does not extend into the margin.

#### Alternative CSS box Model

You might think it is rather inconvenient to have to add up the border and padding to get the real size of the box, and you would be right! For this reason, CSS had an alternative box model introduced some time after the standard box model. Using this model, any width is the width of the visible box on the page, therefore the content area width is that width minus the width for the padding and border. The same CSS as used above would give the below result (width = 350px, height = 150px).

<p align="center">
    <img src="imgs/alternate-box-model.png" width=80%>
</p>

By default, browsers use the standard box model. If you want to turn on the alternative model for an element, you do so by setting ```box-sizing: border-box``` on it. By doing this, you are telling the browser to use the border box, as shown above, as your defined area.

```css
.box {
  box-sizing: border-box;
}
```

If you want all of your elements to use the alternative box model, and this is a common choice among developers, set the box-sizing property on the ```<html>``` element, then set all other elements to inherit that value, as seen in the snippet below. 

```css
html {
  box-sizing: border-box;
}
*, *::before, *::after {
  box-sizing: inherit;
}
```

#### Margins, padding, and borders

You've already seen the ```margin```, ```padding```, and ```border``` properties at work in the example above. The properties used in that example are shorthands and allow us to set all four sides of the box at once. These shorthands also have equivalent longhand properties, which allow control over the different sides of the box individually.

##### Margins

The margin is an invisible space around your box. It pushes other elements away from the box. Margins can have positive or negative values. Setting a negative margin on one side of your box can cause it to overlap other things on the page. Whether you are using the standard or alternative box model, the margin is always added after the size of the visible box has been calculated.

We can control all margins of an element at once using the margin property, or each side individually using the equivalent longhand properties:

- ```margin-top```
- ```margin-right```
- ```margin-bottom```
- ```margin-left```

##### Borders

The border is drawn between the margin and the padding of a box. If you are using the standard box model, the size of the border is added to the ```width``` and ```height``` of the box. If you are using the alternative box model then the size of the border makes the content box smaller as it takes up some of that available ```width``` and ```height```.

For styling borders, there are a large number of properties — there are four borders, and each border has a style, width, and color that we might want to manipulate.

You can set the width, style, or color of all four borders at once using the ``border`` property.

To set the properties of each side individually, you can use:

- ```border-top```
- ```border-right```
- ```border-bottom```
- ```border-left```

To set the width, style, or color of all sides, use the following:

- ```border-width```
- ```border-style```
- ```border-color```

To set the width, style, or color of a single side, you can use one of the more granular longhand properties:

- ```border-top-width```
- ```border-top-style```
- ```border-top-color```
- ```border-right-width```
- ```border-right-style```
- ```border-right-color```
- ```border-bottom-width```
- ```border-bottom-style```
- ```border-bottom-color```
- ```border-left-width```
- ```border-left-style```
- ```border-left-color```

##### Padding

The padding sits between the border and the content area. Unlike margins, you cannot have negative amounts of padding, so the value must be 0 or a positive value. Padding is typically used to push the content away from the border. Any background applied to your element will display behind the padding.

We can control the padding on all sides of an element using the ```padding``` property, or on each side individually using the equivalent longhand properties:

- ```padding-top```
- ```padding-right```
- ```padding-bottom```
- ```padding-left```

#### Using display: inline-block

There is a special value of ```display```, which provides a middle ground between ```inline``` and ```block```. This is useful for situations where you do not want an item to break onto a new line, but do want it to respect ```width``` and ```height``` and avoid the overlapping seen above.

An element with display: inline-block does a subset of the block things we already know about:

- The ```width``` and ```height``` properties are respected.
- ```padding```, ```margin```, and ```border``` will cause other elements to be pushed away from the box.

It does not, however, break onto a new line, and will only become larger than its content if you explicitly add ```width``` and ```height``` properties.

### CSS: Backgrounds and Borders

The CSS ```background``` property is a shorthand for a number of background longhand properties that we will meet in this lesson. If you discover a complex background property in a stylesheet, it might seem a little hard to understand as so many values can be passed in at once.

#### Background colors

The ```background-color``` property defines the background color on any element in CSS. The property accepts any valid ```<color>```. A ```background-color``` extends underneath the content and padding box of the element.

#### Background images

The ```background-image``` property enables the display of an image in the background of an element. In the example below, we have two boxes — one has a background image which is larger than the box (balloons.jpg), the other has a small image of a single star (star.png).

##### Controlling background-repeat
The ```background-repeat``` property is used to control the tiling behavior of images. The available values are:

- ```no-repeat``` — stop the background from repeating altogether.
- ```repeat-x``` — repeat horizontally.
- ```repeat-y``` — repeat vertically.
- ```repeat``` — the default; repeat in both directions.

##### Sizing the background image

The balloons.jpg image used in the initial background images example, is a large image that was cropped due to being larger than the element it is a background of. In this case we could use the ```background-size``` property, which can take length or percentage values, to size the image to fit inside the background.

You can also use keywords:

__```cover```__ — the browser will make the image just large enough so that it completely covers the box area while still retaining its aspect ratio. In this case, part of the image is likely to end up outside the box.

__```contain```__ — the browser will make the image the right size to fit inside the box. In this case, you may end up with gaps on either side or on the top and bottom of the image, if the aspect ratio of the image is different from that of the box.

##### Positioning the background image

The ```background-position``` property allows you to choose the position in which the background image appears on the box it is applied to. This uses a coordinate system in which the top-left-hand corner of the box is ```(0,0)```, and the box is positioned along the horizontal (```x```) and vertical (```y```) axes.

Note: The default ```background-position``` value is ```(0,0)```.

The most common ```background-position``` values take two individual values — a horizontal value followed by a vertical value.

You can use keywords such as ```top``` and ```right``` (look up the others on the ```background-position``` page):

```css
.box {
  background-image: url(star.png);
  background-repeat: no-repeat;
  background-position: top center;
}
```

And Lengths, and percentages:

```css
.box {
  background-image: url(star.png);
  background-repeat: no-repeat;
  background-position: 20px 10%;
}
```

You can also mix keyword values with lengths or percentages, in which case the first value must refer to the horizontal position or offset and the second vertical. For example:

```css
.box {
  background-image: url(star.png);
  background-repeat: no-repeat;
  background-position: 20px top;
}
```

Finally, you can also use a 4-value syntax in order to indicate a distance from certain edges of the box — the length unit, in this case, is an offset from the value that precedes it. So in the CSS below we are positioning the background 20px from the top and 10px from the right:

```css
.box {
  background-image: url(star.png);
  background-repeat: no-repeat;
  background-position: top 20px right 10px;
}
```

##### Gradient backgrounds
A gradient — when used for a background — acts just like an image and is also set by using the ```background-image``` property.

You can read more about the different types of gradients and things you can do with them on the MDN page for the ````<gradient>```` data type. A fun way to play with gradients is to use one of the many CSS Gradient Generators available on the web, such as <a href="https://cssgradient.io/">this one<a>. You can create a gradient then copy and paste out the source code that generates it.

##### Multiple background images
It is also possible to have multiple background images — you specify multiple ```background-image``` values in a single property value, separating each one with a comma.

When you do this you may end up with background images overlapping each other. The backgrounds will layer with the last listed background image at the bottom of the stack, and each previous image stacking on top of the one that follows it in the code.

> Note: Gradients can be happily mixed with regular background images.

The other ```background-*``` properties can also have comma-separated values in the same way as ```background-image```:

```css
background-image: url(image1.png), url(image2.png), url(image3.png), url(image4.png);
background-repeat: no-repeat, repeat-x, repeat;
background-position: 10px 20px,  top right;
```

Each value of the different properties will match up to the values in the same position in the other properties. Above, for example, ```image1```'s ```background-repeat``` value will be ```no-repeat```. However, what happens when different properties have different numbers of values? The answer is that the smaller numbers of values will cycle — in the above example there are four background images but only two ```background-position``` values. The first two position values will be applied to the first two images, then they will cycle back around again — ```image3``` will be given the first position value, and ```image4``` will be given the second position value.

##### Background attachment
Another option we have available for backgrounds is specifying how they scroll when the content scrolls. This is controlled using the ```background-attachment``` property, which can take the following values:

- ```scroll```: causes the element's background to scroll when the page is scrolled. If the element content is scrolled, the background does not move. In effect, the background is fixed to the same position on the page, so it scrolls as the page scrolls.
- ```fixed```: causes an element's background to be fixed to the viewport so that it doesn't scroll when the page or element content is scrolled. It will always remain in the same position on the screen.
- ```local```: fixes the background to the element it is set on, so when you scroll the element, the background scrolls with it.
The ```background-attachment``` property only has an effect when there is content to scroll, so we've made a demo to demonstrate the differences between the three values — have a look at background-attachment.html (also see the source code here).

##### Using the background shorthand property
As I mentioned at the beginning of this lesson, you will often see backgrounds specified using the ```background``` property. This shorthand lets you set all of the different properties at once.

If using multiple backgrounds, you need to specify all of the properties for the first background, then add your next background after a comma. In the example below we have a gradient with a size and position, then an image background with no-repeat and a position, then a color.

There are a few rules that need to be followed when writing background image shorthand values, for example:

- A ```background-color``` may only be specified after the final comma.
- The value of ```background-size``` may only be included immediately after ```background-position```, separated with the '/' character, like this:``` center/80%```.

Take a look at the MDN page for ```background``` to see all of the considerations.

#### Borders
When learning about the Box Model, we discovered how borders affect the size of our box. In this lesson we will look at how to use borders creatively. Typically when we add borders to an element with CSS we use a shorthand property that sets the ```color```, ```width```, and ```style``` of the border in one line of CSS.

We can set a border for all four sides of a box with ```border```:

```css
.box {
  border: 1px solid black;
}
```

Or we can target one edge of the box, for example:

```css
.box {
  border-top: 1px solid black;
}
```

The individual properties for these shorthands would be:

```css
.box {
  border-width: 1px;
  border-style: solid;
  border-color: black;
}
```

And for the longhands:

```css
.box {
  border-top-width: 1px;
  border-top-style: solid;
  border-top-color: black;
}
```

> Note: These top, right, bottom, and left border properties also have mapped logical properties that relate to the writing mode of the document (e.g. left-to-right or right-to-left text, or top-to-bottom). We'll be exploring these in the next lesson, which covers handling different text directions.

##### Borders: Rounded corners
Rounding corners on a box is achieved by using the ```border-radius``` property and associated longhands which relate to each corner of the box. Two lengths or percentages may be used as a value, the first value defining the horizontal radius, and the second the vertical radius. In a lot of cases, you will only pass in one value, which will be used for both.

For example, to make all four corners of a box have a 10px radius:

```css
.box {
  border-radius: 10px;
}
```

Or to make the top right corner have a horizontal radius of 1em, and a vertical radius of 10%:

```css
.box {
  border-top-right-radius: 1em 10%;
}
```

### CSS: Handling Different Text Directions

https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Handling_different_text_directions

### CSS: Handling Overflowing Content

https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Overflowing_content

#### The overflow property
The  ```overflow``` property is how you take control of an element's overflow. It is the way you instruct the browser how it should behave. The default value of overflow is visible. With this default, we can see content when it overflows.

To crop content when it overflows, you can set overflow: hidden. This does exactly what it says: it hides overflow. Beware that this can make some content invisible. You should only do this if hiding content won't cause problems.

### CSS: Values and Units

Every property used in CSS has a value type defining the set of values that are allowed for that property.  Taking a look at any property page on MDN will help you understand the values associated with a value type that are valid for any particular property. In this lesson we will take a look at some of the most frequently used value types, and their most common values and units.

In CSS specifications and on the property pages here on MDN you will be able to spot value types as they will be surrounded by angle brackets, such as ```<color>``` or ```<length>```. When you see the value type ```<color>``` as valid for a particular property, that means you can use any valid color as a value for that property, as listed on the ```<color>``` reference page.

> Note: You'll also see CSS value types referred to as data types. The terms are basically interchangeable — when you see something in CSS referred to as a data type, it is really just a fancy way of saying value type. The term value refers to any particular expression supported by a value type that you choose to use.

> Note: Yes, CSS value types tend to be denoted using angle brackets to differentiate them from CSS properties (e.g., the color property, versus the ```<color>``` data type). You might get confused between CSS data types and HTML elements too, as they both use angle brackets, but this is unlikely — they are used in very different contexts.

In the following example, we have set the color of our heading using a keyword, and the background using the rgb() function:

```css
h1 {
  color: black;
  background-color: rgb(197,93,161);
}
```

A value type in CSS is a way to define a collection of allowable values. This means that if you see ```<color>``` as valid you don't need to wonder which of the different types of color value can be used — keywords, hex values, rgb() functions, etc. You can use any available ```<color>``` values, assuming they are supported by your browser. The page on MDN for each value will give you information about browser support. For example, if you look at the page for ```<color>``` you will see that the browser compatibility section lists different types of color values and support for them.

#### Numbers, lengths, and percentages
There are various numeric value types that you might find yourself using in CSS. The following are all classed as numeric:

<table class="standard-table no-markdown">
  <thead>
    <tr>
      <th scope="col">Data type</th>
      <th scope="col">Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code><a href="/en-US/docs/Web/CSS/integer">&lt;integer&gt;</a></code></td>
      <td>
        An <code>&lt;integer&gt;</code> is a whole number such as
        <code>1024</code> or <code>-55</code>.
      </td>
    </tr>
    <tr>
      <td><code><a href="/en-US/docs/Web/CSS/number">&lt;number&gt;</a></code></td>
      <td>
        A <code>&lt;number&gt;</code> represents a decimal number — it may or may
        not have a decimal point with a fractional component. For
        example,&nbsp;<code>0.255</code>, <code>128</code>, or <code>-1.2</code>.
      </td>
    </tr>
    <tr>
      <td><code><a href="/en-US/docs/Web/CSS/dimension">&lt;dimension&gt;</a></code></td>
      <td>
        A <code>&lt;dimension&gt;</code> is a <code>&lt;number&gt;</code> with a
        unit attached to it. For example,&nbsp;<code>45deg</code>, <code>5s</code>,
        or <code>10px</code>. <code>&lt;dimension&gt;</code> is an umbrella
        category that includes the
        <code><a href="/en-US/docs/Web/CSS/length">&lt;length&gt;</a></code>, <code><a href="/en-US/docs/Web/CSS/angle">&lt;angle&gt;</a></code>, <code><a href="/en-US/docs/Web/CSS/time">&lt;time&gt;</a></code>, and
        <code><a href="/en-US/docs/Web/CSS/resolution">&lt;resolution&gt;</a></code>
        types.
      </td>
    </tr>
    <tr>
      <td><code><a href="/en-US/docs/Web/CSS/percentage">&lt;percentage&gt;</a></code></td>
      <td>
        A <code>&lt;percentage&gt;</code> represents a fraction of some other
        value. For example,&nbsp;<code>50%</code>. Percentage values are always
        relative to another quantity. For example, an element's length is
        relative to its parent element's length.
      </td>
    </tr>
  </tbody>
</table>

##### Lengths
The numeric type you will come across most frequently is <length>. For example 10px (pixels) or 30em. There are two types of lengths used in CSS — relative and absolute. It's important to know the difference in order to understand how big things will become.

##### Absolute length units
The following are all absolute length units — they are not relative to anything else, and are generally considered to always be the same size.

<table>
  <thead>
    <tr>
      <th>Unit</th>
      <th>Name</th>
      <th>Equivalent to</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>cm</code></td>
      <td>Centimeters</td>
      <td>1cm = 37.8px = 25.2/64in</td>
    </tr>
    <tr>
      <td><code>mm</code></td>
      <td>Millimeters</td>
      <td>1mm = 1/10th of 1cm</td>
    </tr>
    <tr>
      <td><code>Q</code></td>
      <td>Quarter-millimeters</td>
      <td>1Q = 1/40th of 1cm</td>
    </tr>
    <tr>
      <td><code>in</code></td>
      <td>Inches</td>
      <td>1in = 2.54cm = 96px</td>
    </tr>
    <tr>
      <td><code>pc</code></td>
      <td>Picas</td>
      <td>1pc = 1/6th of 1in</td>
    </tr>
    <tr>
      <td><code>pt</code></td>
      <td>Points</td>
      <td>1pt = 1/72th of 1in</td>
    </tr>
    <tr>
      <td><code>px</code></td>
      <td>Pixels</td>
      <td>1px = 1/96th of 1in</td>
    </tr>
  </tbody>
</table>

##### Relative length units
Relative length units are relative to something else, perhaps the size of the parent element's font, or the size of the viewport. The benefit of using relative units is that with some careful planning you can make it so the size of text or other elements scales relative to everything else on the page. Some of the most useful units for web development are listed in the table below.

<table class="standard-table no-markdown">
  <thead>
    <tr>
      <th scope="col">Unit</th>
      <th scope="col">Relative to</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>em</code></td>
      <td>
        Font size of the parent, in the case of typographical properties like
        <code><a href="/en-US/docs/Web/CSS/font-size">font-size</a></code>, and font size of the element itself, in the case of other properties
        like <code><a href="/en-US/docs/Web/CSS/width">width</a></code>.
      </td>
    </tr>
    <tr>
      <td><code>ex</code></td>
      <td>x-height of the element's font.</td>
    </tr>
    <tr>
      <td><code>ch</code></td>
      <td>The advance measure (width) of the glyph "0" of the element's font.</td>
    </tr>
    <tr>
      <td><code>rem</code></td>
      <td>Font size of the root element.</td>
    </tr>
    <tr>
      <td><code>lh</code></td>
      <td>Line height of the element.</td>
    </tr>
    <tr>
      <td><code>vw</code></td>
      <td>1% of the viewport's width.</td>
    </tr>
    <tr>
      <td><code>vh</code></td>
      <td>1% of the viewport's height.</td>
    </tr>
    <tr>
      <td><code>vmin</code></td>
      <td>1% of the viewport's smaller dimension.</td>
    </tr>
    <tr>
      <td><code>vmax</code></td>
      <td>1% of the viewport's larger dimension.</td>
    </tr>
  </tbody>
</table>

##### ```ems``` and ```rems```
```em``` and ```rem``` are the two relative lengths you are likely to encounter most frequently when sizing anything from boxes to text. It's worth understanding how these work, and the differences between them, especially when you start getting on to more complex subjects like styling text or CSS layout. The below example provides a demonstration.

The HTML illustrated below is a set of nested lists — we have three lists in total and both examples have the same HTML. The only difference is that the first has a class of ems and the second a class of rems.

To start with, we set 16px as the font size on the ```<html>``` element.

To recap, the em unit means "my parent element's font-size" in the case of typography. The ```<li>``` elements inside the ```<ul>``` with a class of ems take their sizing from their parent. So each successive level of nesting gets progressively larger, as each has its font size set to 1.3em — 1.3 times its parent's font size.

To recap, the rem unit means "The root element's font-size". (rem stands for "root em".) The ```<li>``` elements inside the ```<ul>``` with a class of rems take their sizing from the root element (```<html>```). This means that each successive level of nesting does not keep getting larger.

However, if you change the ```<html>``` font-size in the CSS you will see that everything else changes relative to it — both rem- and em-sized text.

##### Percentages
In a lot of cases, a percentage is treated in the same way as a length. The thing with percentages is that they are always set relative to some other value. For example, if you set an element's ```font-size``` as a percentage, it will be a percentage of the ```font-size``` of the element's parent. If you use a percentage for a width value, it will be a percentage of the width of the parent.

In the below example the two percentage-sized boxes and the two pixel-sized boxes have the same class names. The sets are 40% and 200px wide respectively.

The difference is that the second set of two boxes is inside a wrapper that is 400 pixels wide. The second 200px wide box is the same width as the first one, but the second 40% box is now 40% of 400px — a lot narrower than the first one!

##### Numbers
Some value types accept numbers, without any unit added to them. An example of a property which accepts a unitless number is the opacity property, which controls the opacity of an element (how transparent it is). This property accepts a number between 0 (fully transparent) and 1 (fully opaque).

##### Color
There are many ways to specify color in CSS, some of which are more recently implemented than others. The same color values can be used everywhere in CSS, whether you are specifying text color, background color, or whatever else.

The standard color system available in modern computers is 24 bit, which allows the display of about 16.7 million distinct colors via a combination of different red, green and blue channels with 256 different values per channel (256 x 256 x 256 = 16,777,216.) Let's have a look at some of the ways in which we can specify colors in CSS.

##### Color keywords
Quite often in examples here in the learn section or elsewhere on MDN you will see the color keywords used, as they are a simple and understandable way of specifying color. There are a number of these keywords, some of which have fairly entertaining names! You can see a full list on the page for the ```<color>``` value type.

##### Hexadecimal RGB values
The next type of color value you are likely to encounter is hexadecimal codes. Each hex value consists of a hash/pound symbol (#) followed by six hexadecimal numbers, each of which can take one of 16 values between 0 and f (which represents 15) — so 0123456789abcdef. Each pair of values represents one of the channels — red, green and blue — and allows us to specify any of the 256 available values for each (16 x 16 = 256.)

These values are a bit more complex and less easy to understand, but they are a lot more versatile than keywords — you can use hex values to represent any color you want to use in your color scheme.

##### RGB and RGBA values
The third scheme we'll talk about here is RGB. An RGB value is a function — rgb() — which is given three parameters that represent the red, green, and blue channel values of the colors, in much the same way as hex values. The difference with RGB is that each channel is represented not by two hex digits, but by a decimal number between 0 and 255 — somewhat easier to understand.

You can also use RGBA colors — these work in exactly the same way as RGB colors, and so you can use any RGB values. However, there is a fourth value that represents the alpha channel of the color, which controls opacity. If you set this value to 0 it will make the color fully transparent, whereas 1 will make it fully opaque. Values in between give you different levels of transparency.

> Note: Setting an alpha channel on a color has one key difference to using the ```opacity``` property we looked at earlier. When you use opacity you make the element and everything inside it opaque, whereas using RGBA colors only makes the color you are specifying opaque.

##### HSL and HSLA values
Slightly less well-supported than RGB is the HSL color model (not supported on old versions of IE), which was implemented after much interest from designers. Instead of red, green, and blue values, the hsl() function accepts hue, saturation, and lightness values, which are used to distinguish between the 16.7 million colors, but in a different way:

- ```Hue```: The base shade of the color. This takes a value between 0 and 360, representing the angles around a color wheel.
- ```Saturation```: How saturated is the color? This takes a value from 0–100%, where 0 is no color (it will appear as a shade of grey), and 100% is full color saturation
- ```Lightness```: How light or bright is the color? This takes a value from 0–100%, where 0 is no light (it will appear completely black) and 100% is full light (it will appear completely white)

##### Images
The ```<image>``` value type is used wherever an image is a valid value. This can be an actual image file pointed to via a url() function, or a gradient.

Note: there are some other possible values for ```<image>```, however these are newer and currently have poor browser support. Check out the page on MDN for the ```<image>``` data type if you want to read about them.

Position
The ```<position>``` value type represents a set of 2D coordinates, used to position an item such as a background image (via background-position). It can take keywords such as top, left, bottom, right, and center to align items with specific bounds of a 2D box, along with lengths, which represent offsets from the top and left-hand edges of the box.

A typical position value consists of two values — the first sets the position horizontally, the second vertically. If you only specify values for one axis the other will default to center.

##### Strings and identifiers

Throughout the examples above, we've seen places where keywords are used as a value (for example ```<color>``` keywords like red, black, rebeccapurple, and goldenrod). These keywords are more accurately described as identifiers, a special value that CSS understands. As such they are not quoted — they are not treated as strings.

##### Functions
The final type of value we will take a look at is the group of values known as functions. In programming, a function is a reusable section of code that can be run multiple times to complete a repetitive task with minimum effort on the part of both the developer and the computer. Functions are usually associated with languages like JavaScript, Python, or C++, but they do exist in CSS too, as property values. We've already seen functions in action in the Colors section — ```rgb()```, ```hsl()```, etc. The value used to return an image from a file — ```url()``` — is also a function.

A value that behaves more like something you might find in a traditional programming language is the ```calc()``` CSS function. This function gives you the ability to do simple calculations inside your CSS. It's particularly useful if you want to work out values that you can't define when writing the CSS for your project, and need the browser to work out for you at runtime.

### CSS: Sizing Items in CSS

In the various lessons so far, you have come across a number of ways to size items on a web page using CSS. Understanding how big the different features in your design will be is important. So, in this lesson we will summarize the various ways elements get a size via CSS and define a few terms about sizing that will help you in the future.

#### Natural, Intrinsic Size of Things
HTML Elements have a natural size, set before they are affected by any CSS. A straightforward example is an image. An image file contains sizing information, described as its __intrinsic size__. This size is determined by the image itself, not by any formatting we happen to apply.

If you place an image on a page and do not change its height or width, either by using attributes on the <img> tag or else by CSS, it will be displayed using that intrinsic size.

#### Setting a specific size
We can, of course, give elements in our design a specific size. When a size is given to an element (the content of which then needs to fit into that size) we refer to it as an extrinsic size. Take our ```<div>``` from the example above — we can give it specific ```width``` and ```height``` values, and it will now have that size no matter what content is placed into it. As we discovered in our previous lesson on overflow, <u>a set height can cause content to overflow if there is more content than the element has space to fit inside it</u>.

Due to this problem of overflow, fixing the height of elements with lengths or percentages is something we need to do very carefully on the web.

#### Using percentages
In many ways, percentages act like length units, and as we discussed in the lesson on values and units, they can often be used interchangeably with lengths. When using a percentage you need to be aware what it is a percentage of. In the case of a box inside another container, if you give the child box a percentage width it will be a percentage of the width of the parent container.

```css
.box {
  border: 5px solid darkblue;
  width: 50%;
}
```
```html
<div class="box">
  I have a percentage width.
</div>
```

This is because percentages resolve against the size of the containing block. With no percentage applied our ```<div>``` would take up 100% of the available space, as it is a block level element. If we give it a percentage width, this becomes a percentage of the space it would normally fill.

##### Percentage margins and padding
If you set margins and padding as a percentage, you may notice some strange behavior.

You might expect for example the percentage top and bottom margins to be a percentage of the element's height, and the percentage left and right margins to be a percentage of the element's width. However, this is not the case!

When you use margin and padding set in percentages, the value is calculated from the inline size of the containing block — therefore the width when working in a horizontal language. In our example, all of the margins and padding are 10% of the width. This means you can have equal-sized margins and padding all around the box. This is a fact worth remembering if you do use percentages in this way.

##### min- and max- sizes
In addition to giving things a fixed size, we can ask CSS to give an element a minimum or a maximum size. If you have a box that might contain a variable amount of content, and you always want it to be at least a certain height, you could set the min-height property on it. The box will always be at least this height, but will then grow taller if there is more content than the box has space for at its minimum height.

This is very useful for dealing with variable amounts of content while avoiding overflow.

A common use of ```max-width``` is to cause images to scale down if there is not enough space to display them at their intrinsic width while making sure they don't become larger than that width.

As an example, if you were to set ```width: 100%``` on an image, and its intrinsic width was smaller than its container, the image would be forced to stretch and become larger, causing it to look pixelated.

If you instead use ```max-width: 100%```, and its intrinsic width is smaller than its container, the image will not be forced to stretch and become larger, thus preventing pixelation.

In the example below, we have used the same image three times. The first image has been given width: 100% and is in a container which is larger than it, therefore it stretches to the container width. The second image has ```max-width: 100%``` set on it and therefore does not stretch to fill the container. The third box contains the same image again, also with ```max-width: 100%``` set; in this case you can see how it has scaled down to fit into the box.

This technique is used to make images _responsive_, so that when viewed on a smaller device they scale down appropriately. You should, however, not use this technique to load really large images and then scale them down in the browser. Images should be appropriately sized to be no larger than they need to be for the largest size they are displayed in the design. Downloading overly large images will cause your site to become slow, and it can cost users more money if they are on a metered connection.

##### Viewport units
The viewport — which is the visible area of your page in the browser you are using to view a site — also has a size. In CSS we have units which relate to the size of the viewport — the ```vw``` unit for viewport width, and ```vh``` for viewport height. Using these units you can size something relative to the viewport of the user.

```1vh``` is equal to ```1%``` of the viewport height, and ```1vw``` is equal to ```1%``` of the viewport width. You can use these units to size boxes, but also text. In the example below we have a box which is sized as ```20vh``` and ```20vw```. The box contains a letter A, which has been given a ```font-size``` of ```10vh```.

### CSS: Images, Media, and Form Elements

In this lesson we will take a look at how certain special elements are treated in CSS. Images, other media, and form elements behave a little differently from regular boxes in terms of your ability to style them with CSS. Understanding what is and isn't possible can save some frustration, and this lesson will highlight some of the main things that you need to know.

#### Replaced elements
Images and video are described as replaced elements. This means that CSS cannot affect the internal layout of these elements — only their position on the page amongst other elements. As we will see however, there are various things that CSS can do with an image.

Certain replaced elements, such as images and video, are also described as having an aspect ratio. This means that it has a size in both the horizontal (x) and vertical (y) dimensions, and will be displayed using the intrinsic dimensions of the file by default.

#### Sizing images
As you already know from following these lessons, everything in CSS generates a box. If you place an image inside a box that is smaller or larger than the intrinsic dimensions of the image file in either direction, it will either appear smaller than the box, or overflow the box. You need to make a decision about what happens with the overflow.

So what can we do about the overflowing issue?

As we learned in our previous lesson, a common technique is to make the max-width of an image 100%. This will enable the image to become smaller in size than the box but not larger. This technique will also work with other replaced elements such as ```<video>```s, or ```<iframe>```s.

Try adding ```max-width: 100%``` to the ```<img>``` element in the example above. You will see that the smaller image remains unchanged, but the larger one becomes smaller to fit into the box.

You can make other choices about images inside containers. For example, you may want to size an image so it completely covers a box.

The ```object-fit``` property can help you here. When using ```object-fit``` the replaced element can be sized to fit a box in a variety of ways.

Below we have used the value cover, which sizes the image down, maintaining the aspect ratio so that it neatly fills the box. As the aspect ratio is maintained, some parts of the image will be cropped by the box.

```css
.box {
  width: 200px;
  height: 200px;
}

img {
  height: 100%;
  width: 100%;
}

.cover {
  object-fit: cover;
}

.contain {
  object-fit: contain;
}
```
```html
<div class="wrapper">
  <div class="box"><img src="balloons.jpg" alt="balloons" class="cover"></div>
  <div class="box"><img src="balloons.jpg" alt="balloons" class="contain"></div>
</div>
```

If we use ```contain``` as a value, the image will be scaled down until it is small enough to fit inside the box. This will result in "letterboxing" if it is not the same aspect ratio as the box.

You could also try the value of ```fill```, which will fill the box but not maintain the aspect ratio.

#### Replaced elements in layout
When using various CSS layout techniques on replaced elements, you may well find that they behave slightly differently from other elements. For example, in a flex or grid layout elements are stretched by default to fill the entire area. Images will not stretch, and instead will be aligned to the start of the grid area or flex container.

If you are following these lessons in order then you may not have looked at layout yet. Just keep in mind that replaced elements, when they become part of a grid or flex layout, have different default behaviors, essentially to avoid them being stretched strangely by the layout.

To force the image to stretch to fill the grid cell it is in, you'd have to do something like the following:

```css
img {
  width: 100%;
  height: 100%;
}
```

### CSS: Form Elements

Form elements can be a tricky issue when it comes to styling with CSS. The Web Forms module contains detailed guides to the trickier aspects of styling these, which I will not fully reproduce here. There are, however, a few key basics worth highlighting in this section.

Many form controls are added to your page by way of the ```<input>``` element — this defines simple form fields such as text inputs, through to more complex fields added in HTML5 such as color and date pickers. There are some additional elements, such as ```<textarea>``` for multiline text input, and also elements used to contain and label parts of forms such as ```<fieldset>``` and ```<legend>```.

HTML5 also contains attributes that enable web developers to indicate which fields are required, and even the type of content that needs to be entered. If the user enters something unexpected, or leaves a required field blank, the browser can show an error message. Different browsers vary with one another in how much styling and customization they allow for such items.

#### Styling text input elements
Elements that allow for text input, such as ```<input type="text">```, and the more specific ```<input type="email">```, and the ```<textarea>``` element are quite easy to style and tend to behave just like other boxes on your page. The default styling of these elements will differ, however, based on the operating system and browser that your user visits the site with.

In the example below we have styled some text inputs using CSS — you can see that things such as borders, margins and padding all apply as you would expect. We are using attribute selectors to target the different input types. Try changing how this form looks by adjusting the borders, adding background colors to the fields, and changing fonts and padding.

```css
input[type="text"],
input[type="email"] {
  border: 2px solid #000;
  margin: 0 0 1em 0;
  padding: 10px;
  width: 100%;
}

input[type="submit"] {
  border: 3px solid #333;
  background-color: #999;
  border-radius: 5px;
  padding: 10px 2em;
  font-weight: bold;
  color: #fff;
}

input[type="submit"]:hover {
  background-color: #333;
}
```
```html
<form>
  <div><label for="name">Name</label>
  <input type="text" id="name"></div>
  <div><label for="email">Email</label>
  <input type="email" id="email"></div>

  <div class="buttons"><input type="submit" value="Submit"></div>
</form>
```

> Warning: You should take care when changing the styling of form elements to make sure it is still obvious to the user they are form elements. You could create a form input with no borders and background that is almost indistinguishable from the content around it, but this would make it very hard to recognise and fill in.

As explained in the lessons on form styling in the HTML part of this course, many of the more complex input types are rendered by the operating system and are inaccessible to styling. You should therefore always assume that forms are going to look quite different for different visitors and test complex forms in a number of browsers.

#### Inheritance and form elements
In some browsers, form elements do not inherit font styling by default. Therefore, if you want to be sure that your form fields use the font defined on the body, or on a parent element, you should add this rule to your CSS.

```css
button,
input,
select,
textarea {
  font-family : inherit;
  font-size : 100%;
}
```

#### Form elements and box-sizing
Across browsers, form elements use different box sizing rules for different widgets. You learned about the ```box-sizing``` property in our box model lesson and you can use this knowledge when styling forms to ensure a consistent experience when setting widths and heights on form elements.

For consistency, it is a good idea to set margins and padding to 0 on all elements, then add these back in when styling particular controls:

```css
button,
input,
select,
textarea {
  box-sizing: border-box;
  padding: 0;
  margin: 0;
}
```

#### Other useful settings
In addition to the rules mentioned above, you should also set ```overflow: auto``` on ```<textarea>```s to stop IE showing a scrollbar when there is no need for one:

```css
textarea {
  overflow: auto;
}
```

#### Putting it all together into a "reset"
As a final step, we can wrap up the various properties discussed above into the following "form reset" to provide a consistent base to work from. This includes all the items mentioned in the last three sections:

```css
button,
input,
select,
textarea {
  font-family: inherit;
  font-size: 100%;
  box-sizing: border-box;
  padding: 0;
  margin: 0;
}

textarea {
  overflow: auto;
}
```

> Note: Normalizing stylesheets are used by many developers to create a set of baseline styles to use on all projects. Typically these do similar things to those described above, making sure that anything different across browsers is set to a consistent default before you do your own work on the CSS. They are not as important as they once were, as browsers are typically more consistent than in the past. However if you want to take a look at one example, check out <a href="https://necolas.github.io/normalize.css/">Normalize.css<a>, which is a very popular stylesheet used as a base by many projects.

### CSS: Styling Tables

https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Styling_tables

### CSS: Debugging

https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Debugging_CSS

#### Solving specificity issues
Sometimes during development, but in particular when you need to edit the CSS on an existing site, you will find yourself having a hard time getting some CSS to apply. No matter what you do, the element just doesn't seem to take the CSS. What is generally happening here is that a more specific selector is overriding your changes, and here DevTools will really help you out.

In our example file there are two words that have been wrapped in an ```<em>``` element. One is displaying as orange and the other hotpink. In the CSS we have applied:

```css
em {
  color: hotpink;
  font-weight: bold;
}
```

Above that in the stylesheet however is a rule with a ```.special``` selector:

```css
.special {
  color: orange;
}
```

As you will recall from the lesson on cascade and inheritance where we discussed specificity, class selectors are more specific than element selectors, and so this is the value that applies. DevTools can help you find such issues, especially if the information is buried somewhere in a huge stylesheet.

Inspect the ```<em>``` with the class of .special and DevTools will show you that orange is the color that applies, and also that the color property applied to the ```<em>``` is crossed out. You can now see that the class selector is overriding the element selector.

### CSS: Organizing CSS

https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Organizing

A good tip is to add a block of comments between logical sections in your stylesheet too, to help locate different sections quickly when scanning it, or even to give you something to search for to jump right into that part of the CSS. If you use a string that won't appear in the code, you can jump from section to section by searching for it — below we have used ||.

```css
/* || General styles */

...

/* || Typography */

...

/* || Header and Main Navigation */

...
```

#### CSS methodologies
Instead of needing to come up with your own rules for writing CSS, you may benefit from adopting one of the approaches already designed by the community and tested across many projects. These methodologies are essentially CSS coding guides that take a very structured approach to writing and organizing CSS. Typically they tend to render CSS more verbosely than you might have if you wrote and optimised every selector to a custom set of rules for that project.

However, you do gain a lot of structure by adopting one. Since many of these systems are widely used, other developers are more likely to understand the approach you are using and be able to write their own CSS in the same way, rather than having to work out your own personal methodology from scratch.

#### OOCSS
Most of the approaches you will encounter owe something to the concept of Object Oriented CSS (OOCSS), an approach made popular by the work of Nicole Sullivan. The basic idea of OOCSS is to separate your CSS into reusable objects, which can be used anywhere you need on your site. The standard example of OOCSS is the pattern described as The Media Object. This is a pattern with a fixed size image, video or other element on one side, and flexible content on the other. It's a pattern we see all over websites for comments, listings, and so on.

If you are not taking an OOCSS approach you might create a custom CSS for the different places this pattern is used, for example, by creating two classes, one called comment with a bunch of rules for the component parts, and another called list-item with almost the same rules as the comment class except for some tiny differences. The differences between these two components are the list-item has a bottom border, and images in comments have a border whereas list-item images do not.

```css
.comment {
  display: grid;
  grid-template-columns: 1fr 3fr;
}

.comment img {
  border: 1px solid grey;
}

.comment .content {
  font-size: .8rem;
}

.list-item {
  display: grid;
  grid-template-columns: 1fr 3fr;
  border-bottom: 1px solid grey;
}

.list-item .content {
  font-size: .8rem;
}
```

In OOCSS, you would create one pattern called ```media``` that would have all of the common CSS for both patterns — a base class for things that are generally the shape of the media object. Then we'd add an additional class to deal with those tiny differences, thus extending that styling in specific ways.

```css
.media {
  display: grid;
  grid-template-columns: 1fr 3fr;
}

.media .content {
  font-size: .8rem;
}

.comment img {
  border: 1px solid grey;
}

.list-item {
  border-bottom: 1px solid grey;
}
```

In your HTML, the comment would need both the media and comment classes applied:

```html
<div class="media comment">
  <img />
  <div class="content"></div>
</div>
```

The list-item would have media and list-item applied:

```html
<ul>
  <li class="media list-item">
    <img />
    <div class="content"></div>
  </li>
</ul>
```

## CSS styling text

https://developer.mozilla.org/en-US/docs/Learn/CSS/Styling_text

## CSS layout

https://developer.mozilla.org/en-US/docs/Learn/CSS/CSS_layout/Introduction
https://developer.mozilla.org/en-US/docs/Learn/CSS/CSS_layout

# Lec 2: CSS

```css
/* selectors */

/* know different selector types: type selector, class, id */
/* know how to combine selectors: " " to combine, "," to apply to more than one thing */
```

### CSS Layout

> last week: linear flow. style objects. float images. selectors.

This week, we will learn how to control the division of elements to make things more fancy. 

### CSS Columns
https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Columns

> This works like text in a news paper: text arranged into columns.

```css
/* have text move horizontally in this many columns.*/
/* for everything class "container" */
.container {
  column-count: 3;
}
```

```html
<div class="container">
  <p>Veggies es bonus vobis, proinde vos postulo essum magis kohlrabi welsh onion daikon amaranth tatsoi tomatillo melon azuki bean garlic.</p>

  <p>Gumbo beet greens corn soko endive gumbo gourd. Parsley shallot courgette tatsoi pea sprouts fava bean collard greens dandelion okra wakame tomato. Dandelion cucumber earthnut pea peanut soko zucchini.</p>
            
  <p>Turnip greens yarrow ricebean rutabaga endive cauliflower sea lettuce kohlrabi amaranth water spinach avocado daikon napa cabbage asparagus winter purslane kale. Celery potato scallion desert raisin horseradish spinach carrot soko.</p>
</div>
```


__```column-span```__ attribute allows an element to span across all columns when its value is set to ```all```
```css
/* object with id: spanningelement should span accross anything its in*/
#spanningelement {
  column-span: all;
}
```

example
```html
<article>
  <h2>Header spanning all of the columns</h2>
  <p>
     The h2 should span all the columns. The rest
     of the text should be distributed among the columns.
  </p>
  <p>This is a bunch of text split into three columns using the CSS `columns` property. The text is equally distributed over the columns.</p>
  <p>This is a bunch of text split into three columns using the CSS `columns` property. The text is equally distributed over the columns.</p>
  <p>This is a bunch of text split into three columns using the CSS `columns` property. The text is equally distributed over the columns.</p>
  <p>This is a bunch of text split into three columns using the CSS `columns` property. The text is equally distributed over the columns.</p>
</article>
```

```css
article {
  columns: 3;
}

h2 {
  column-span: all;
}
```


For more, see: W3Schools CSS Multiple Columns: https://www.w3schools.com/css/css3_multiple_columns.asp

#### CSS Columns: basic

The simplest layout method uses the ```column-count``` CSS attribute to spread content over multiple columns. You can specify the spacing and borders between column elements, and allow elements like headers to span multiple columns (as below)

```html
<html>
<head>
  <title>Story Pal</title>
  <style>
    .container {
      width: 960px;
      margin: auto;
    }

    .columns {
      column-count: 3;
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>Story Pal</h1>

    <div class="columns">
    <p>Veggies es bonus vobis, proinde vos postulo essum magis kohlrabi welsh onion daikon amaranth tatsoi tomatillo melon azuki bean garlic.</p>

    <p>Gumbo beet greens corn soko endive gumbo gourd. Parsley shallot courgette tatsoi pea sprouts fava bean collard greens dandelion okra wakame tomato. Dandelion cucumber earthnut pea peanut soko zucchini.</p>

    <p>Turnip greens yarrow ricebean rutabaga endive cauliflower sea lettuce kohlrabi amaranth water spinach avocado daikon napa cabbage asparagus winter purslane kale. Celery potato scallion desert raisin horseradish spinach carrot soko.</p>
    </div>
  </div>
</body>

```
#### CSS Columns: Masonry Layout

If you want columns to be roughly packed together, not needing to be aligned horizontally, use the ```CSS Columns - Masonry Layout```. Good for packing things of various sizes in a compact space. Various assests of different sizes.

https://www.smashingmagazine.com/2019/01/css-multiple-column-layout-multicol/#masonry-like-display-of-content

### CSS Flexbox

__```Flexbox```__ is a one-dimensional layout method for arranging items in rows or columns. Items flex (expand) to fill additional space or shrink to fit into smaller spaces.

> Paradigm became popular and part of the standard CSS standards.

You can do nice simple things with it and intricate things with it. Powerful. For a long time, this was the only way to do cross browser layouts were with floats and positioning. These worked, but it was frustrating. ```Flexbox``` made this easier to do. With ```Flexbox```, you can

1. Vertically center a block of content inside its parent.
2. Make all the children of a container take up an equal amount of the available width/height, regardless of how much width/height is available.
3. Make all columns in a multiple-column layout adopt the same height even if they contain a different amount of content.

If we want to make a bunch of elements lay out as flexible boxes, we call upon their parent element and edit the __```display```__ attribute:

```css
section {
  display: flex;
}
```

This causes the ```section``` element to become a __flex container__ and its children to become __flex items__. The result of this should be something like so:

<p align="center">
    <img src="imgs/flexbox.png" width=100%>
</p>

> __Note__ also that you can use a display value of __```inline-flex```__ if you wish to lay out an element's children as flex items, but have that element behave like an inline element.

<p align="center">
    <img src="imgs/flex_terms.png" width=100%>
</p>

Terms
- __main axis__: axis running in the direction the flex items are laid out in (e.g. rows across the page, or columns down the page). The start, end are called __main start__, __main end__
- __cross axis__: axis running perpendicular to the direction the flex items are laid out in. The start, end are called __cross start__, __cross end__
- __flex container__: parent element that has ```display: flex``` (above, it is ```<section>```).
- __flex items__: items laid out as flexible boxes inside the flex container (above, it is the ```<article>``` elements)

Change direction of the flexing with __```flex-direction```__. By default, it is set to row
```css
section {
  display: flex;
  /* flex-direction: row;...by default */
}

section {
  display: flex;
  flex-direction: column;
}
```

You can reverse the column or row ordering with __```row-reverse```__ and __```column-reverse```__.

example:
```html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Flexbox 0 — starting code</title>
    <style>
      html {
        font-family: sans-serif;
      }

      body {
        margin: 0;
      }

      header {
        background: purple;
        height: 100px;
      }

      h1 {
        text-align: center;
        color: white;
        line-height: 100px;
        margin: 0;
      }

      article {
        padding: 10px;
        margin: 10px;
        background: aqua;
      }

      /* Add your flexbox CSS below here */
      section {
        display: flex;
      }
     
    </style>
  </head>
  <body>
    <header>
      <h1>Sample flexbox example</h1>
    </header>

    <section>
      <article>
        <h2>First article</h2>

        <p>Tacos actually microdosing, pour-over semiotics banjo chicharrones retro fanny pack portland everyday carry vinyl typewriter. Tacos PBR&B pork belly, everyday carry ennui pickled sriracha normcore hashtag polaroid single-origin coffee cold-pressed. PBR&B tattooed trust fund twee, leggings salvia iPhone photo booth health goth gastropub hammock.</p>
      </article>

      <article>
        <h2>Second article</h2>

        <p>Tacos actually microdosing, pour-over semiotics banjo chicharrones retro fanny pack portland everyday carry vinyl typewriter. Tacos PBR&B pork belly, everyday carry ennui pickled sriracha normcore hashtag polaroid single-origin coffee cold-pressed. PBR&B tattooed trust fund twee, leggings salvia iPhone photo booth health goth gastropub hammock.</p>
      </article>

      <article>
        <h2>Third article</h2>

        <p>Tacos actually microdosing, pour-over semiotics banjo chicharrones retro fanny pack portland everyday carry vinyl typewriter. Tacos PBR&B pork belly, everyday carry ennui pickled sriracha normcore hashtag polaroid single-origin coffee cold-pressed. PBR&B tattooed trust fund twee, leggings salvia iPhone photo booth health goth gastropub hammock.</p>

        <p>Cray food truck brunch, XOXO +1 keffiyeh pickled chambray waistcoat ennui. Organic small batch paleo 8-bit. Intelligentsia umami wayfarers pickled, asymmetrical kombucha letterpress kitsch leggings cold-pressed squid chartreuse put a bird on it. Listicle pickled man bun cornhole heirloom art party.</p>
      </article>
    </section>
  </body>
</html>
```

#### Flexbox: widths, wrapping, flexible sizing, relative sizes

```css
/* change default flex direction*/
flex-direction: column;
/* can reverse direction*/
row-reverse: true;
column-reverse: true;
```

Can turn on wrapping with __```flex-wrap```__
``` css
/* By default, new lines are not made. */
/* But you can turn on wrapping */
flex-wrap: wrap;
/* specify that each article should be 200px wide */
flex: 200px;
/* specify a min width for column before wrapping */
```

We can control what proportion of the space flex items take up compared to other flex items.

```css
/* Have each article try to take up 1 relative unit of space */
article {
  flex: 1;
}
/* This is a unitless proportion value that dictates how much 
available space along the main axis each flex item will take 
up compared to other flex items. In this case, we're giving 
each <article> element the same value (a value of 1), which 
means they'll all take up an equal amount of the spare space
left after properties like padding and margin have been set.*/

/* have the 3rd article take up twice the relative space of other
articles.*/
article:nth-of-type(3) {
  flex: 2;
}
```

You can specify a minimum size value with the flex value.

```css
article {
  flex: 1 200px;
}

article:nth-of-type(3) {
  flex: 2 200px;
}
/* This basically states, "Each flex item will first be given 
200px of the available space. After that, the rest of the 
available space will be shared according to the proportion 
units." Try refreshing and you'll see a difference in how the
space is shared.
```

__flex shorthand, longhand__

```flex``` is a shorthand property that can specify up to three different values:

- The unitless proportion value we discussed above. This can be specified separately using the ```flex-grow``` longhand property.
- A second unitless proportion value, __```flex-shrink```__, which comes into play when the flex items are overflowing their container. This value specifies how much an item will shrink in order to prevent overflow. This is quite an advanced flexbox feature and we won't be covering it any further in this article.
- The minimum size value we discussed above. This can be specified separately using the ```flex-basis``` longhand value.

We'd advise against using the longhand flex properties unless you really have to (for example, to override something previously set). They lead to a lot of extra code being written, and they can be somewhat confusing.

__Horizontal, Vertical Alignment__

You can use flexbox features to align flex items along the main or cross axis. E.g. making a neat, flexible button/toolbar. Normally, these buttons will jam into the top left-hand corner.

Adding the following horizontally and vertically centers the buttons
```css
div {
  display: flex;
  align-items: center;
  justify-content: space-around;
}
```
__```align-items```__ controls where the flex items sit on the cross axis.

- By default, the value is ```stretch```, which stretches all flex items to fill the parent in the direction of the cross axis. If the parent doesn't have a fixed width in the cross axis direction, then all flex items will become as long as the longest flex item. This is how our first example had columns of equal height by default.
- The ```center``` value that we used in our above code causes the items to maintain their intrinsic dimensions, but be centered along the cross axis. This is why our current example's buttons are centered vertically.
- You can also have values like ```flex-start``` and ```flex-end```, which will align all items at the start and end of the cross axis respectively. See <a href="https://developer.mozilla.org/en-US/docs/Web/CSS/align-items">```align-items```</a> for the full details.

You can override the ```align-items``` behavior for individual flex items by applying the ```align-self``` property to them.
```css
button:first-child {
  align-self: flex-end;
}
```

__```justify-content```__ controls where the flex items sit on the main axis.

- The default value is ```flex-start```, which makes all the items sit at the start of the main axis.
- You can use ```flex-end``` to make them sit at the end.
- ```center``` is also a value for ```justify-content```. It'll make the flex items sit in the center of the main axis.
- The value we've used above, ```space-around```, is useful — it distributes all the items evenly along the main axis with a bit of space left at either end.
There is another value, ```space-between```, which is very similar to ```space-around``` except that it doesn't leave any space at either end.

__Ordering Flex Items__

Flexbox also has a feature for changing the layout order of flex items without affecting the source order.

```css
button:first-child {
  order: 1;
}
```

- By default, all flex items have an ```order``` value of ```0```.
- Flex items with higher specified ```order``` values will appear later in the display order than items with lower order values.
- Flex items with the same ```order``` value will appear in their source order. So if you have four items whose order values have been set as 2, 1, 1, and 0 respectively, their display order would be 4th, 2nd, 3rd, then 1st.
- The 3rd item appears after the 2nd because it has the same order value and is after it in the source order.
- You can set negative ```order``` values to make items appear earlier than items whose value is 0. For example, you could make the "Blush" button appear at the start of the main axis using the following rule:

```css
button:last-child {
  order: -1;
}
```

__Nested Flex Boxes__

It's perfectly OK to set a flex item to also be a flex container, so that its children are also laid out like flexible boxes.

Imagine a layout like so:
```html
section - article
          article
          article - div - button
                    div   button
                    div   button
                          button
                          button
```

```css
/* First, set the children of section to be laid out as flexible boxes */
section {
  display: flex;
}

/* set flex values on the articles themselves. make sure to set the third article
to have its children laid out like flex items too, but this time we lay them
out as columns */

/* all articles same relative size, at least 200 px each */
article {
  flex: 1 200px;
}

/* third article we want to be a flex container, 3x as large. flex flow being column */
article:nth-of-type(3) {
  flex: 3 200px;
  display: flex;
  flex-flow: column;
}

/* select the first div of the 3rd article. set its relative size to 1 100px.
set it to be a flex container (have its button children be laid out like flex items.
lay them out in a wrapping row and align them in the center of the available space*/
article:nth-of-type(3) div:first-child {
  flex: 1 100px;
  display: flex;
  flex-flow: row wrap;
  align-items: center;
  justify-content: space-around;
}

/* Finally, we set some sizing on the button. This time by giving it a flex value
of 1 auto. This has a very interesting effect, which you'll see if you try resizing
your browser window width. The buttons will take up as much space as they can. 
As many will fit on a line as is comfortable; beyond that, they'll drop to a new line.*/
button {
  flex: 1 auto;
  margin: 5px;
  font-size: 18px;
  line-height: 1.5;
}
```

> Idea: write html for simple organizational layout. Then style with CSS. Flexbox can help keep things aligned, e.g. horizontally, with some wrapping rules, styling.

However, you can't do everything with it. Introducing ```CSS grid```

### CSS Grid

__```CSS grid``` helps align things vertically AND horizontally__. Introduces a 2D grid system to CSS. Can be used to lay out major page areas or small user interface elements.

CSS grid enables the following features:
1. __Fiex and flexible track sizes__: can greate a grid with fixed track sizes (e.g. using pixels. sets the grid to the specified pixel which fits to the layout you desire. Can create a grid using flexible sizes with percentages or with the new ```fr``` unit designed for this purpose)
2. __Item placement__: Can place items into a precise location on the grid using line numbers, names, or by targeting an area of the grid. Grid contains an algorithms to control the placement of items not given an explicit position on the grid.
3. __Creation of additional tracks to hold content__: Grid Layout is flexible enough to add additional rows and columns when needed. features included like "add as many columns that will fit into a container"
4. __Alignment control__: Grid contains alignment features so you can control how the items align once placed intoa grid area, and how the entire grid is aligned.
5. __Control of overlapping content__: more than one item can be placed into a grid cell or area, and they can overlap. This layering may then be controlled with the ```z-index``` property.

CSS Grid says: "this is going to be a straight up grid. Here's element1, I want it to start in the top left at (0,0), take up this much space." Can even have elements overlap

#### Grid container

create a __grid container__ with __```display: grid```__, or __```display: inline-grid```__. After, all _direct children_ of that element become _grid items_.

example.
```html
<div class="wrapper">
  <div>One</div>
  <div>Two</div>
  <div>Three</div>
  <div>Four</div>
  <div>Five</div>
</div>
```

```css
.wrapper {
  display: grid;
}
```

#### Grid tracks

define rows and columns on grid with __```grid-template-columns```__, or __```grid-template-rows```__.

```html
<div class="wrapper">
  <div>One</div>
  <div>Two</div>
  <div>Three</div>
  <div>Four</div>
  <div>Five</div>
</div>
```

```css
.wrapper {
  grid-template-columns: 200px, 200px, 200px;
  /* grid-template-columns: 2fr, 1fr, 1fr, */
}
```

<p align="center">
    <img src="imgs/grid_ex.jpg" width=100%>
</p>


Tracks can be defined using any length unit. Grid also includes an additional length unit to create flexible grid tracks.
__```fr```__ represents a fraction of the avialable space in the grid container.

```html
<div class="wrapper">
  <div>One</div>
  <div>Two</div>
  <div>Three</div>
  <div>Four</div>
  <div>Five</div>
</div>
```

```css
.wrapper {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
}
```

<p align="center">
    <img src="imgs/grid_ex2.jpg" width=100%>
</p>

```css
/* can change the relative sizes of each column */
.wrapper {
  display: grid;
  grid-template-columns: 2fr 1fr 1fr;
}
```

Can mix and match specifying absolute and relative sizes. 
```css
/* mixing flexible and absolute sizes

The first track is 500 pixels, so the fixed width is taken
away from the available space. The remaining space is divided
into three and assigned in proportion to the two flexible tracks.*/

.wrapper {
  display: grid;
  grid-template-columns: 500px, 1fr, 1fr,
}
```

When creating the grids above we defined our column tracks with the ```grid-template-columns``` property, but the grid also created rows on its own. These rows are part of the __implicit grid__. Whereas the __explicit grid__ consists of any rows and columns defined with ```grid-template-columns``` or ```grid-template-rows```.

<u>If you place something outside of the defined grid—or due to the amount of content, more grid tracks are needed—then the grid creates rows and columns in the implicit grid</u>. These tracks will be auto-sized by default, resulting in their size being based on the content that is inside them.

You can also define a set size for tracks created in the implicit grid with the ```grid-auto-rows``` and ```grid-auto-columns``` properties.

In the below example we use ```grid-auto-rows``` to ensure that tracks created in the implicit grid are 200 pixels tall.

```css
.wrapper {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  grid-auto-rows: 200px;
}
```

<p align="center">
    <img src="imgs/grid_ex3.JPG" width=100%>
</p>

__Track sizing and minmax__

When setting up an explicit grid or defining the sizing for automatically created rows or columns we may want to give tracks a minimum size, but also ensure they expand to fit any content that is added. _For example, I may want my rows to never collapse smaller than 100 pixels, but if my content stretches to 300 pixels in height, then I would like the row to stretch to that height_.

Use ```minmax()``` function. Below, we use ```minmax()``` in the value of ```grid-auto-rows```. This means automatically created rows will be a minimum of 100 pixels tall, and a maximum of ```auto```. Using ```auto``` means that the size will look at the content size and will stretch to give space for the tallest item in a cell, in this row.

### Grid lines

It should be noted that when we define a grid we define the grid tracks, not the lines. Grid then gives us numbered lines to use when positioning items. In our three column, two row grid we have four column lines.

<p align="center">
    <img src="imgs/grid_ex4.JPG" width=80%>
</p>

__Positioning items against lines__

When placing an item, we target the line – rather than the track.

Example below: place the first two items on our three column track grid, using the ```grid-column-start```, ```grid-column-end```, ```grid-row-start``` and ```grid-row-end``` properties. Working from left to right, the first item is placed against column line 1, and spans to column line 4, which in our case is the far-right line on the grid. It begins at row line 1 and ends at row line 3, therefore spanning two row tracks.

The second item starts on grid column line 1, and spans one track. This is the default so I do not need to specify the end line. It also spans two row tracks from row line 3 to row line 5. The other items will place themselves into empty spaces on the grid.

```html
<div class="wrapper">
  <div class="box1">One</div>
  <div class="box2">Two</div>
  <div class="box3">Three</div>
  <div class="box4">Four</div>
  <div class="box5">Five</div>
</div>
```

```css
.wrapper {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-auto-rows: 100px;
}

.box1 {
  grid-column-start: 1;
  grid-column-end: 4;
  grid-row-start: 1;
  grid-row-end: 3;
}

.box2 {
  grid-column-start: 1;
  grid-row-start: 3;
  grid-row-end: 5;
}
```

<p align="center">
    <img src="imgs/grid_ex5.JPG" width=80%>
</p>

__Line-positioning shorthands__

The longhand values used above can be compressed onto one line for columns with grid-column, and one line for rows with grid-row. 

The following example would give the same positioning as in the previous code, but with far less CSS. The value before the forward slash character (```/```) is the start line, the value after the end line.

```css
.wrapper {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  grid-auto-rows: 100px;
}

.box1 {
  grid-column: 1 / 4;
  grid-row: 1 / 3;
}

.box2 {
  grid-column: 1;
  grid-row: 3 / 5;
}
```

__```grid cell```__: Smallest unit on a grid.

__```grid area```__: When an item spans one or more cells by both row or by column, it creates a _grid area_

__Gutters__ or __alleys__ between grid cells can be created using the ```column-gap``` and ```row-gap``` properties, or the shorthand gap

```html
<div class="wrapper">
  <div>One</div>
  <div>Two</div>
  <div>Three</div>
  <div>Four</div>
  <div>Five</div>
</div>
```

```css
.wrapper {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  column-gap: 10px;
  row-gap: 1em;
}
```

<p align="center">
    <img src="imgs/grid_ex6.JPG" width=80%>
</p>

### CSS Grid: Nesting

https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Grid_Layout/Basic_Concepts_of_Grid_Layout#nesting_grids

### CSS Grid: Layering items with ```z-index```

https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Grid_Layout/Basic_Concepts_of_Grid_Layout#layering_items_with_z-index


So, with Grid we can control number of columns, set absolute/relative widths and sizes. It's like flexbox, but they can stay locked. __Only want to use the grid if you have to__. E.g. use it to layout the basic website structure. Then have segments have Flexbox can be more self-contained. 

> How I develop: start with chrome. drag my browser to different sizes to see if i like what happens (first pass). If doing commercial production of websites for large releases, you will probably try it on many browsers and phones.

### Flexbox vs Grid

Do I only need to control the layout by row __or__ column?: __use flexbox__

Do I only need to control the layout by row __AND__ column?: __use grid__

Flexbox: define container. Fill it. It detemrines how everything should fill (and wrap). The size of the overall flexbox is determined by the number of elements inside the flexbox. You might want the items to determine the size of the website (then flexbox makes sense). However, if designing an online newspaper, you'll use grid to create a defined structure and then fill it with flexible things (create a macro grid and fill it with flex boxes)

In addition to the one-dimensional versus two-dimensional distinction, there is another way to decide if you should use flexbox or grid for a layout. Flexbox works from the content out. An ideal use case for flexbox is when you have a set of items and want to space them out evenly in a container. You let the size of the content decide how much individual space each item takes up. If the items wrap onto a new line, they will work out their spacing based on their size and the available space on that line.

Grid works from the layout in. When you use CSS Grid Layout you create a layout and then you place items into it, or you allow the auto-placement rules to place the items into the grid cells according to that strict grid. It is possible to create tracks that respond to the size of the content, however, they will also change the entire track.

_See Uchicago Website_: https://www.uchicago.edu/

### Advanced Styling: Icons

- [Web Icons Tutorial (W3Schools)](https://www.w3schools.com/icons/)
- [Font Awesome](https://fontawesome.com/)
- [Material Design (Google)](https://materializecss.com/icons.html)
- [Custom Icon Fonts (Fontello)](http://fontello.com/)

with icons and .svg files you can resize them and generate more pixels (e.g. generated by math). Icons can get treated as text and allows you to be able to style the icons like text. You can align them with letters. Dropping images on a line with text, it is hard to make the images re-size to the size of text.

Icons are helpful to put on buttons, headers, etc.

Use like so:
```html
<!DOCTYPE html>
<html>
<head>
    <title>Icons Example</title>
    <!-- add link to a style sheet that loads a font -->
    <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons" />
    <style>
        i {
            color: red;
        }
    </style>
    </head>
<body>

    <!-- <i> is the icon element. stylesheet replaces the words with icons e.g. "cloud" becomes an image -->
    <!-- icons are technically fonts -->
    <p><i class="material-icons">cloud</i>Coulds!</p>
    <p><i class="material-icons">favorite</i></p>
    <i class="material-icons">attachment</i>
    <i class="material-icons">computer</i>
    <i class="material-icons">traffic</i>

</body>
</html>
```

### Responsive Design

- [Media Queries (Mozilla Documentation)](https://developer.mozilla.org/en-US/docs/Web/CSS/Media_Queries/Using_media_queries)
- [Responsive Design (Mozilla Documentation)](https://developer.mozilla.org/en-US/docs/Learn/CSS/CSS_layout/Responsive_Design)

#### Media Queries

see documentation on media queries
> apply these rules if the statement in the media query is true

```css
/* */
@media print { ... }
/* */
@media screen, print { ... }

/* if screen width less than some value apply rules*/
@media (max-width: 12450px) { ... }
/* if screen width less than some value apply rules*/
@media (min-width: 12450px) { ... }
```

> write defualt style sheet. Then, later down the sheet, define alternative stylesheets

Can apply different styling rules under different conditions

### Advanced Styling: Web Fonts

- [Web Fonts (Mozilla Documentation)](https://developer.mozilla.org/en-US/docs/Learn/CSS/Styling_text/Web_fonts)
- [Google Fonts](https://fonts.google.com/)

To get what you want to show on a user's machine if they don't have your font installed:

provide a list of comma separated values for fonts:

```css
p {
  font-family: Helvetica, "Trbuchet MS", Verdana, sans-serif;
}
```

Can have a ruleset at the start of teh CSS, which specifies custom font to load

```css
@font-face {
  font-family: "myFont";
  src: url("myFont.Woff2");
}
```

## Accessibility (a11y)

- [Accessibility (Mozilla Documentation)](https://developer.mozilla.org/en-US/docs/Web/Accessibility)
- [Handling Common Accessibility Problems (Mozilla Documentation)](https://developer.mozilla.org/en-US/docs/Learn/Tools_and_testing/Cross_browser_testing/Accessibility)
  - Screen reader compatibility
  - Keyboard UI controls
  - Color and contrast
  - Alt text captioning

Refers to physical accessibility and technical accessibility. E.g. ```screen reader compatibility```: people who are blind may use readers that read out loud the text on a page. This is why it's important that the html communicate the semantics of the page (article, header, paragraph, etc...).

```keyboard UI controls```: for people who have difficulty navigating on the screen (e.g. hard to use a mouse, hard to use a track pad). You will want to make sure people can navigate with arrow keys. Allowing "tab" to cycle through elements on the page. Making sure people can use enter button to enter something. Sometimes javascript can prevent tabbing from working. Use semantic markup in the html to help with all these problems

```Color and contrast```: color-blindness (red-green, blue-yellow). choose sets of colors keeping in mind potential color-blindness. Low-contrast text is very hard to read. Have a certain minimum saturation difference between text and background.

```Alt text captioning```: providing text messages that caption images.

> Write semantic html and add alt text captioning

## Advanced Styling: CSS Filters and Animation

- [Linear Gradient](https://developer.mozilla.org/en-US/docs/Web/CSS/linear-gradient)
- [Filters](https://developer.mozilla.org/en-US/docs/Web/CSS/filter)
- [Animation](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Animations/Using_CSS_animations)
- [Parallax Scrolling](https://developers.google.com/web/updates/2016/12/performant-parallaxing)

e.g.
```css
.animate img {
  width: 210px;
  animation-duration: 3s;
  animation-name: rise;
}

@keyframes rise {
  from {
    margin-top: 60%;
    margin-left: 100%
  }
  to {
    margin-top: 0%;
  }
}
```

## CSS Resets and Frameworks

- `box-sizing` [(Mozilla Documentation)](https://developer.mozilla.org/en-US/docs/Web/CSS/box-sizing)
- [Eric Meyer's CSS Reset](https://meyerweb.com/eric/tools/css/reset/)
- [Normalize](https://github.com/necolas/normalize.css/)
- [Bootstrap](https://getbootstrap.com/)

In the early history of the internet, browsers would have different interpretations of how things should be rendered based on instructions. e.g. things would be different sizes in Netscape vs Internet Explorer. People began writing code dedicated to browsers (e.g. if on internet explorer do this; else, do this)

>infamous example: ```box-sizing```

__CSS Reset__:Custom css rules to override the browser default. 

__Normalize__: another example of compatibility

__Bootstrap__: framework meant to unify html/css styling.

## CSS References

[CSS Layout: Evolution Detailed by Example (Medium.com)](https://medium.com/@rebeccapeltz/css-layout-evolution-detailed-by-example-579b31475392)

# Server-Side Rendering with LAMP stack

## HTTP

> For a long time a popular interview question at Google was "what happens when you type www.google.com in your browser". 

**Hypertext Transfer Protocol** (**HTTP**) is an application protocol for
distributed, collaborative, hypermedia information systems.[1] HTTP is the
foundation of data communication for the World Wide Web, where hypertext
documents include hyperlinks to other resources that the user can easily access,
for example by a mouse click or by tapping the screen in a web browser.

([Wikipedia](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol))

[HTTP Overview (Mozilla Documentation)](https://developer.mozilla.org/en-US/docs/Web/HTTP/Overview)

[How Browsers Work](https://www.html5rocks.com/en/tutorials/internals/howbrowserswork/)

### HTTP: Requests

[HTTP Overview - Requests (Mozilla Documentation)](https://developer.mozilla.org/en-US/docs/Web/HTTP/Overview#requests)

[HTTP Request Methods (Mozilla Documentation)](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods)

![HTTP request diagram showing protocol version, status code, status message, and headers](images/http_request.png "HTTP request diagram")

Can view the requests that happen at a webpage, dynamically, with inspector tools -> network.

HTML, CSS, Web APIs, and JavaScript can be transmitted via HTTP. The lower level protocols computers use to talk to each other may not support being able to transmit this information (DNS, UDP, TCP, TLS, IP).

### HTTP Flow

Flow begins with the client (usually the browser but can be anything that sends requests like a command line). Our client is something who wants to do something on the internet, they send a request out to a server, and they get a response back.

```method```: what kind of request you're sending (```GET```). This will vary the most on your requests.

```path```: on the server where I'm sending the request to (e.g. Uchicagowebdev.com) (```/``` the root)

```version of the protocol```: the verson of HTTP you are using (```HTTP/1.1```)

> HTTP/1.1 is the most stable verison

```Headers```: (key, value) pairs, where the key is identified by a short string with no spaces in it, a semi-colon, and then a value after. "when I send my request, I have some meta-data about my request. This might tell the server more things about who I am and what I want." ```Accept-Language: fr``` allows you to specify you want to work with the french version/language of the server, but there may not always be such a version. If a server doesn't know what to do with a header, it will ignore it. 

```
GET / HTTP/1.1
Host: developer.mozilla.org
Accept-Language: fr
```

### HTTP Request Methods

__```Get```__: is a common method. This request is idempotent, meaning you will get the same response for every request. This means the request will never change the server.

__```Post```__: is a common method. This request submits an entity to the specified resource, often causing a change in state of side effects on the server.

__```Put```__: less common. Modify in-place (implies pre-existence). Usually 'Post' for new things and 'Put' for things that already exists

__```Delete```__: deletes the specified resource

__```Patch```__: less common. Applies partial modifications to a resource. e.g. if database of addresses, patch would edit a field (attribute) of tuples without affecting the other attributes.

```Post``` usually also contains data that we send to the server

> "I've never seem HEAD, CONNECT used in the wild"

### HTTP Responses

[HTTP Overview - Responses (Mozilla Documentation)](https://developer.mozilla.org/en-US/docs/Web/HTTP/Overview#responses)

[HTTP Response Status Codes (Mozilla Documentation)](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)

![HTTP response diagram showing protocol version, status code, status message, and headers](images/http_response.png "HTTP response diagram")

```version of the protocol```: the verson of HTTP you are using (```HTTP/1.1```)

```status code```: a code that tells you success/failure of a request (```200```)

> hundreds place on response code tells you what general type of response you're getting.
tens and ones place tell you specific information about how it went. 2 hundreds signify success.

```status message```: a message that tells you success/failure of a request (```OK```)

```Headers```: 0 or more headers

```content of response```: actual response from the server in html

```
HTTP/1.1 200 OK
Date: Sat, 09 Oct 2010 14:28:02 GMT
Server: Apache
Last-Modified: ...
...
```

#### Status Codes

> __Successful Responses__. ```200 Success```: is most common success code. 

> __Redirection Messages__. You'll rarely see ```300 Multiple Choice``` responses. That's the server saying
"I didn't find what you wanted so I'm redirecting you elsewhere." ```304 Not Modified```: you requested from the server
something you already have and nothing has changed. So I'm redirecting you to what you already have in your cache.
```302 Found```: "this page has moved." You'll see this if you go to the root of a page, and it redirects you somewhere else.
There are two kinds of redirects. ```302``` means "hey you thought this thing was here, but it's been moved. Keep asking here 
in the future though". You'll see this a lot if you have content that moves around or you change your URL scheme, and it's been changed temporarily.
```301 Moved Permanently```: never ask from here again. E.g. you moved your whole website somewhere else. Use the ```301``` if you are for sure going somewhere else and you never want the browser to go back. E.g. you are permanently changing your domain name. Important because search engines that crawl the web that will respect the 301. So, be very sure something has permenantly migrated to use this. e.g. in my day job, we are getting a new domain name, we will issue 301 when move to new site.

> __Client Error Responses__. ```400 Bad Request```: server could not understand the request due to invalid syntax. Most common is ```404 Not Found```: Server can not find the requested resource. ```403 Forbidden```: client does not have access rights to the content; this requires password and you don't have one. ```401 Unauthorized```: basically, you're not logged in (must be authenticated first). ```429 Too Many Requests```: common when dealing with API's. Machine way of saying "slow down, you've been making too many requests". Usually implmemented with a exponential cool-down timer. You might rate-limit requests because your server might be only so fast, and if you let clients call too many requests the server may not be able to keep up.

> __Server Error Responses__. Means something went wrong on server. I'm just part of a server that's giving you web responses, I don't know exactly what went wrong, but something bad happened. ```500 Internal Server Error```: you'll get some of this in the process of working on assignments. Might happen if a server divides by 0, or throws an exception. The exception gets bubbled up to the web server, and then it says "well I can't send this exception to the user (becuase it might expose information about how the server works thus exposing security risks), so I'll just throw a 500 error." ```503 Service Unavailable``` might be used when a server is down for maintenance. A common thing we will do is parse out the status code, and then conditionally do different things depending on the status code. e.g. if 400 error, display error to user, if 200, display results to user.

#### ```Curl```

Command line tool for sending, receiving HTTP requests

```bash
curl -i uchicagowebdev.com
```

>"means send an HTTP request to uchicagowebdev.com, and the default values are a GET request." response is just text.

```
HTTP/1.1 200 OK
Date: Wed, 26 Jan 2022 16:04:50 GMT
Server: Apache/2.4.51 ()
Upgrade: h2,h2c
Connection: Upgrade
Last-Modified: Fri, 14 Jan 2022 22:47:02 GMT
ETag: "1ec-5d59294d84f58"
Accept-Ranges: bytes
Content-Length: 492
Content-Type: text/html; charset=UTF-8

<html>
<head>
<title>MPCS 52553 | Web Development</title>
</head>
<body>
<h1>Hello Class!</h1>
<hr/>
<p><a href="helloclass/">Student pages</a></p>
<p><a href="https://github.com/UChicagoWebDev">Course GitHub</a></p>
<p><a href="https://classroom.github.com/classrooms/97002355-uchicago-web-development-winter-2022-classroom-1">GitHub Classroom</a> (for submitting assignments)</p>
<p><a href="https://app.slack.com/client/T71CT0472/C02TBJ5BHU2">Course Slack</a></p>
<style>

</body>
</html>
```

> We got a 200 response. great. We got some headers (key: value pairse). ```Content-Length``` can be useful. Tells you / the browser how long the content will be, indicating when we're done reading a response. ```Content-Type``` tells you what kind of content you are seeing. you / the browser is receiving zeros and ones; ```charset```: determines how to display / interpret those zeros and ones. ```UTF-8``` is a pretty universal character set. Then, we get the content. This is the file: index/html. The browser knows "ok this is a body. I'm going to take that and render it on the page"

Without the ```-i``` flag, It's just going to give us the body; it will strip the body out. If I run curl with ```-v``` for verbose, it will also show the get request.

```
curl -v uchicagowebdev.com
```

## Web Servers: Apache

> We've talked about HTTP requests and responses. Client sends HTTP requests to servers and receive responses back. We've learned a client can be many things: your web browser, your command line (the curl utility). Apache is free. Now, we're going to learn about the server.

> The Apache HTTP Server ("httpd") was launched in 1995 and it has been the most
popular web server on the Internet since April 1996. It has celebrated its 20th
birthday as a project in February 2015. ([httpd.apache.org](https://httpd.apache.org/)) The web server running at http://uchicagowebdev.com/ is Apache. By default, when it receives an HTTP request, it looks for files on the local filesystem that match
the URL path and returns them to the web browser.

__Apache__: released 1995. Very popular open-source web server software. It's been the most popular web server for over 25 years. We use an Apache server for our uchicagowebdev.com. How Apache works by default: you run it, there's a specific directory inside our web server that Apache looks at. When you send a request to an Apache server, the root path is that specific directory that it's looking for, and then if there is a index.html, it will serve that page. If you pass it subsequent paths, it looks for directories with that path, it looks for files in those directories with that path. If they exist in the file system on the server, then Apache will send a 200 response and then as the body just include the contents of that file. You're streaming the file contents back to the client as the body of your http request. If I try to go to a URL where this is not a file on the filesystem, it will return a 404. If there's no html found, Apache will display its splash page.

use ```cat``` to view contents of a file in bash.

Anything that accepts HTTP requests and sends HTTP responses is a web server.

> This completes our discussion on webpages and being able to host that webpage to display to user.

### Server-Side Rendering with PHP

Remember: a web server is a web server if it accepts HTTP requests and replies with HTTP responses. For example, a server could read a file and do things with it, and then send a response back.

For example, there is a file called ```Math.php```. It looks like a standard HTML document to start, but then we have something new. We have a ```<p></p>``` element, and inside that we have ```<?php echo 2+2 ?>```. This is encoding a script. 

```html
<html>
<head>
  <title>PHP Math Examples</title>
  <link rel="stylesheet" type="text/css" href="css/style.css">
</head>
<body>
  <h1>Let's do some math on the server!</h1>
  <p>2 + 2 = <?php echo 2+2 ?></p>
  <p>8 - 3 = <?php echo 8-3 ?></p>
  <p>6 * 7 = <?php echo 6*7 ?></p>
  <p>pi = <?php echo pi()?></p>
</body>
```

The php code is getting run on the server, and the ```?``` is replaced with the result of the expression. Before the server passes math.php to the browser, it evaluates the code in the <?php ?> tags. If you View Source, you'll see that they have been replaced with plain text. Importantly, that evaluation happens on the server, before any data is sent to the user.

Before the server passes math.php to the browser, it evaluates the code in the
`<?php ?>` tags. If you View Source, you'll see that they have been replaced
with plain text. Importantly, that evaluation happens on the server, before any
data is sent to the user.

```html
<html>
<head>
  <title>PHP Math Examples</title>
  <link rel="stylesheet" type="text/css" href="css/style.css">
</head>
<body>
  <h1>Let's do some math on the server!</h1>
  <p>2 + 2 = 4</p>
  <p>8 - 3 = 5</p>
  <p>6 * 7 = 42</p>
  <p>pi = 3.1415926535898></p>
</body>
```

Note, we don't know whether this HTML we received from the server side was from a static webpage or from a dynamic webpage. We see it as what some final HTML document.

#### PHP

PHP is a whole, Turing-complete programming language. Famously, Facebook is written using it. We also use it extensively at my day job.

[PHP](https://www.php.net/) is a whole, Turing-complete programming language.
Famously, Facebook is written using it. We also use it extensively at my day
job.

So a web server isn't just a program that shows existing html documents to
users; it can be any arbitrarily complex program that returns HTTP responses.

![php.net](images/php.png "Screenshot of php.net")

#### LAMP Stack

One very popular thing you might want to do with your program that responds to HTTP requests is have it interact with a database. And for a long time the most popular such configuration of these pieces was the [LAMP](https://en.wikipedia.org/wiki/LAMP_(software_bundle) stack.

> Originally popularized from the phrase "Linux, Apache, MySQL, and PHP", the acronym "LAMP" now refers to a generic software stack model. The modularity of a LAMP stack may vary, but this particular software combination has become popular because it is sufficient to host a wide variety of web site frameworks, such as WordPress. The components of the LAMP stack are present in the software repositories of most Linux distributions. https://en.wikipedia.org/wiki/LAMP_(software_bundle)

__LAMP__: (__L__)inux, (__A__)pache, (__M__)ySQL, (__P__)HP. Linux is free, Apache is free, MySQL is free, Apache is free.

__WordPress__: uses the LAMP stack. It uses Apache, MySQL, PHP. 75% of all traffic on the internet interacted with webservers on the WordPress (a user of the LAMP stack)

#### Lab: Working with Submission Forms

[Dealing with Forms (PHP Documentation)](https://www.php.net/manual/en/tutorial.forms.php)

[http://uchicagowebdev.com/post.php](http://uchicagowebdev.com/post.php)

[http://uchicagowebdev.com/post_no_escape.php](http://uchicagowebdev.com/post_no_escape.php)

Typically, you make a ```form``` html element:

```html
<form action="action.php" method="post">
 <p>Your name: <input type="text" name="name" /></p>
 <p>Your age: <input type="text" name="age" /></p>
 <p><input type="submit" /></p>
</form>
```

Forms have default behavior where they have an ```action```: "when this form is activated, and a ```submit``` button will say if its clicked, it will send an http request and load a new page. Where will it send the request? It's defined in the ```action``` tag. In this case, the server will find ```action.php```, run the script, and then send the response back, and that will be the new page displayed to the user (as if they made a simple GET request).

Inside PHP code, we get access to a special global variable called ```_POST```, which is an associative array of all the inputs that were in our form. So when we send this request, it takes the contents of the input forms and sends them along as part of the request, and on the server side it parses them out and puts them in a special variable called ```_POST``` (in PHP variables are prepended with a ```$``` sign to indicate variable). So below we have an associative array with 2 fields, and the values that the user entered in the form are saved and made available on the server side in this array.

```php
Hi <?php echo htmlspecialchars($_POST['name']); ?>.
You are <?php echo (int)$_POST['age']; ?> years old.
```

can read php file like so
```bash
cat action.php
```

```post.php```. standard html at the top. Inside the ```post-place``` div at the top: "at the time I'm rendering this page, if this came as the result of a post request (because there's a submit action in the ```_POST``` variable array), then what I want you to do is take the input that the user put into our ```freeform``` text area and I want you to print it back out.
```php
<html>
<head>
  <title>POST Handling Example</title>
  <link rel="stylesheet" type="text/css" href="css/style.css">
  <style> hr { margin: 20px 20px 20px 20px; } </style>
</head>
<body>
  <h1>Post Code</h1>
  <hr/>
  <div class="post-place">
    <?php if (isset($_POST['submit'])) {?>
      <?php echo htmlspecialchars($_POST['freeform']) ?>
    <?php } ?>
  </div>
  <hr/>
  <div>
    <form action="post.php" method="post">
      <label for="freeform">Any old text</label>
      <textarea name="freeform"></textarea>
      <input type="submit" name="submit" value="Submit"></input>
    </form>
  </div>
</body>

```

<p align="center">
    <img src="imgs/php_form.jpg" width=100%>
</p>

Any time we take user input, we need to be careful. We run something that clients send us. Need to make sure we do not take an input they give us and end up interpretting it as code. For example someone might write text that could become part of the page source (security liability).

Anytime you take user input and want to display it on the page, wrap it with a function that processes that text for vulnerabilities. It renders everything as text. Use __```htmlspecialchars()```__ function. It is PHP's __escape function__.

The reason this works that we can go to Math.php: there's a default built in setting in Apache that says "when I go to fetch a file from the filesystem to fetch it for the user, if it is a php file, run it through a php interpretter, and send the results of that to the user (can turn this off). If this is turned off, the html will receive the php tags and not know how to deal with them.

> Recap: this paradigm is popular because you can co-mingle php code with html.

> take a look at doing if statements in php. This ```if``` gets evaluated on the server-side. You can use this to handle, for instance, to do password checking. The client would not be able to see this password check. Once we are on the server executing arbitrary code, we may want to interact with a database.

## Databases

[MySQL (Wikipedia)](https://en.wikipedia.org/wiki/MySQL)

[Introduction to Relational Databases (MariaDB Documentation)](https://mariadb.com/kb/en/introduction-to-relational-databases/)

[A MariaDB Primer](https://mariadb.com/kb/en/a-mariadb-primer/)

![Database Tables](images/tables.png "Visualization of two Database Tables")

Relational database: you have a bunch of tables with attributes that have datatypes. Each row is a tuple. You can have columns in one table refer to columns in another table. e.g. table1 is a catalog in a store (Products). This is everything you sell. In a separate table, you can have a database keeping track of who bought each product (Orders). Allows us to reduce data duplication and keep all data in one place.

```MySQL``` is still a very popular go to database. It's very well supported and free. Literal ```MySQL``` is not used much more. The Swedish company that developed ```MySQL``` for free was bought by SunMicrosystems (early manufacturer of internet hardware and internet software). SunMicrosystems got bought by Oracle (wealthy maker of expensive database software). Now, when new versions of MySQL come out, they are developed by Oracle, who have a nasty reputation of chasing people down to make sure they have certain licenses. So, there is now a new open-source version of MySQL called ```MariaDB```.

We will use Apache, LAMP stack, php. However, we learned that setting up MySQL in previous classes is an enormous logistical headache. If running a windows machine, it can work but substantially harder to setup than if using LINUX/mac machine. Instead of using MySQL, we will use an alternative __```SQLite```__.

### SQLite

Has largely the same syntax as ```MySQL``` or ```Postgres```. We will use SQLite instead of MySQL for our 'LAMP' stack.

```
sqlite3                // fires up sqlite (open interpretter)
.databases             // view databases in workspace
.open example_db       // example_db is a file on file system...doing this will create a file if doesn't exist
.tables                // view list of tables in database
.schema example_table  // view schema of a given table
```

```SQLite``` would be good for putting on IOT devices. Good when you need an onboard database for something that does not have a lot of storage. SQLite is written in C, the installation is on the order of hundreds of kilobytes.

#### SELECT, INSERT

```sql
-- select everything from a table
SELECT * from posts;
-- select everything from a table
SELECT * from comments;
/* insert new book into database */
INSERT INTO books (Title, SeriesID, AuthorID)
VALUES ("Lair of Bones",2,2);
```

#### JOIN

[SQL Joins](https://www.w3schools.com/sql/sql_join.asp)

```sql
SELECT Orders.OrderID, Customers.CustomerName, Orders.OrderDate
FROM Orders
INNER JOIN Customers ON Orders.CustomerID=Customers.CustomerID;
```

Rely on the __foreign key__ relationship between tables to ensure that, for instance, when adding an order to an Orders table, that the customer ID specified in the order is a customer ID that actually exists in the Customers database.

Here are the different types of joins in SQL:

__```(INNER) JOIN```__: Returns records that have matching values in both tables

__```LEFT (OUTER) JOIN```__: Returns all records from the left table, and the matched records from the right table

__```RIGHT (OUTER) JOIN```__: Returns all records from the right table, and the matched records from the left table

__```FULL (OUTER) JOIN```__: Returns all records when there is a match in either left or right table

<p align="center">
    <img src="imgs/join_types.jpg" width=100%>
</p>

### Accessing the Database with PHP

[Connecting to the Database](https://www.php.net/manual/en/intro.pdo.php)

[Prepared Statements to Send Queries](https://www.php.net/manual/en/pdo.prepared-statements.php)

> We've alluded to this. While our PHP code is executing, we may want it to connect to a library. Indeed there is a php library that allows you to do this. It's called __```database-specific PDO driver```__. It allows you to write php code that directly interacts with your database. We will be using SQLlite.

We will be connecting to a database on our local machine.

Note, if we allow a user to enter free text, and we send those things to the database, a clever user can use the free-text and break out of the planned query you planned to write and then write their own arbitrary query afterwards. This kind of security flaw is called a __SQL injection__.

```sql
-- example
Robert'); DROP TABLE Students; --?
```

In the same way we used HTML escape to take any HTML special characters and cast them as their display characters instead of parsing them as HTML, what we will do is instead of taking the user entered data and appending it to a string and sending it to the database, we will make a prepared statement. Always use prepared statements. Do not use string manipulation (regex, pattern matching, string appending, etc).

```php
<?php
// instead of just saing: INSERT INTO REGISTRY (name, value) VALUES
// we can say take the insert statement, declare 2 variables it will take (:name and :value)
// we bind the values to our variables with bindParam()
$stmt = $dbh->prepare("INSERT INTO REGISTRY (name, value) VALUES (:name, :value)");
$stmt->bindParam(':name', $name);
$stmt->bindParam(':value', $value);

// insert one row
$name = 'one';
$value = 1;
$stmt->execute();

// insert another row with different values
$name = 'two';
$value = 2;
$stmt->execute();
?>
```

> So, Apache will interpret php as code and run it first before sending back to client. Now we have all the ingredients to make web application. Web page is the interface, we have a web server where we run an application, process information and interact with databases via PHP. This is great because all the software can run on servers, where we do not need to worry about how software will run on client computers. Now, we just need a client to be able to connect to our server. We do not need to worry about shipping. If you can get on the internet, you can run our program.


### Lec 3

#### curl

```bash
curl -i examplecat.com      // 301 error: permanent redirection
```

response
```
$ curl -i http://examplecat.com
HTTP/1.1 301 Moved Permanently
cache-control: public, max-age=0, must-revalidate
content-length: 38
content-type: text/plain
date: Fri, 28 Jan 2022 12:20:29 GMT
location: https://examplecat.com/
server: Netlify
x-nf-request-id: 01FTHH3J5HSKKNDWD19ZKP2CEP
age: 40930
```

Act as a different browser on the client side. For example, you may get a different response depending on the type of client sending the request.
```bash
curl -v https://wikipedia.org -A "Internet Explorer"
```

Can use curl to send post requests. Note code below may not be right. In this case, wikipedia would say "uh, no I can't do anything with that."
```bash
curl -d "This is my data" https://en.wikipedia.org
```

To copy something to a server
```
scp filename.php student@uchicagowebdev.com/var/www/html/helloclass
```
### WordPress

Popular framework that runs in PHP (LAMP stack).

### Setting up WAMP

Follow here:
https://www.youtube.com/watch?v=I2vExxv5JGo
