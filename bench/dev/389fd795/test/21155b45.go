      
    
    
      <h4>Response</h4>
      <pre><code>Status: 200 OK</code></pre>
      <div class="height-constrained-code-block"></div>
    
    
      <h4>Notes</h4>
      <ul class="mt-2">
      
        <li><a href="/en/developers/apps">Works with GitHub Apps</a></li>
      
      </ul>
    
    
  </div>
  <hr>
</div>
<h2 id="feeds"><a href="#feeds">Feeds</a></h2>
  <div>
  <div>
    <h3 id="get-feeds" class="pt-3">
      <a href="#get-feeds">Get feeds</a>
    </h3>
    <p>GitHub provides several timeline resources in <a href="http://en.wikipedia.org/wiki/Atom_(standard)">Atom</a> format. The Feeds API lists all the feeds available to the authenticated user:</p>
<ul>
<li><strong>Timeline</strong>: The GitHub global public timeline</li>
<li><strong>User</strong>: The public timeline for any user, using <a href="https://developer.github.com/v3/#hypermedia">URI template</a></li>
<li><strong>Current user public</strong>: The public timeline for the authenticated user</li>
<li><strong>Current user</strong>: The private timeline for the authenticated user</li>
<li><strong>Current user actor</strong>: The private timeline for activity created by the authenticated user</li>
<li><strong>Current user organizations</strong>: The private timeline for the organizations the authenticated user is a member of.</li>
<li><strong>Security advisories</strong>: A collection of public announcements that provide information about security-related vulnerabilities in software on GitHub.</li>
</ul>
<p><strong>Note</strong>: Private feeds are only returned when <a href="https://developer.github.com/v3/#basic-authentication">authenticating via Basic Auth</a> since current feed URIs use the older, non revocable auth tokens.</p>
  </div>
  <pre><code><span class="bg-blue text-white rounded-1 px-2 py-1" style="text-transform: uppercase">get</span> /feeds</code></pre>
  <div>
    
      <h4 id="get-feeds--parameters">
        <a href="#get-feeds--parameters">Parameters</a>
      </h4>
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Type</th>
            <th>In</th>
            <th>Description</th>
          </tr>
          </h5>
          <pre><code class="hljs language-shell">curl \
  -H &quot;Accept: application/vnd.github.v3+json&quot; \
  https://api.github.com/feeds
</code></pre>
        
      
        
          <h5>
            
              JavaScript (<a href="https://github.com/octokit/core.js#readme">@octokit/core.js</a>)
            
          </h5>
          <pre><code class="hljs language-javascript"><span class="hljs-keyword">await</span> octokit.request(<span class="hljs-string">&apos;GET /feeds&apos;</span>)
</code></pre>
        
      
    
    
      <h4>Default response</h4>
      <pre><code>Status: 200 OK</code></pre>
      <div class="height-constrained-code-block"><pre><code class="hljs language-json">{
  <span class="hljs-attr">&quot;timeline_url&quot;</span>: <span class="hljs-string">&quot;https://github.com/timeline&quot;</span>,
  <span class="hljs-attr">&quot;user_url&quot;</span>: <span class="hljs-string">&quot;https://github.com/{user}&quot;</span>,
  <span class="hljs-attr">&quot;current_user_public_url&quot;</span>: <span class="hljs-string">&quot;https://github.com/octocat&quot;</span>,
  <span class="hljs-attr">&quot;current_user_url&quot;</span>: <span class="hljs-string">&quot;https://github.com/octocat.private?token=abc123&quot;</span>,
  <span class="hljs-attr">&quot;current_user_actor_url&quot;</span>: <span class="hljs-string">&quot;https://github.com/octocat.private.actor?token=abc123&quot;</span>,
  <span class="hljs-attr">&quot;current_user_organization_url&quot;</span>: <span class="hljs-string">&quot;&quot;</span>,
  <span class="hljs-attr">&quot;current_user_organization_urls&quot;</span>: [
    <span class="hljs-string">&quot;https://github.com/organizations/github/octocat.private.atom?token=abc123&quot;</span>
  ],
  <span class="hljs-attr">&quot;security_advisories_url&quot;</span>: <span class="hljs-string">&quot;https://github.com/security-advisories&quot;</span>,
  <span class="hljs-attr">&quot;_links&quot;</span>: {
    <span class="hljs-attr">&quot;timeline&quot;</span>: {
      <span class="hljs-attr">&quot;href&quot;</span>: <span class="hljs-string">&quot;https://github.com/timeline&quot;</span>,
      <span class="hljs-attr">&quot;type&quot;</span>: <span class="hljs-string">&quot;application/atom+xml&quot;</span>
    },
    <span class="hljs-attr">&quot;user&quot;</span>: {
      <span class="hljs-attr">&quot;href&quot;</span>: <span class="hljs-string">&quot;https://github.com/{user}&quot;</span>,
      <span class="hljs-attr">&quot;type&quot;</span>: <span class="hljs-string">&quot;application/atom+xml&quot;</span>
    },
    <span class="hljs-attr">&quot;current_user_public&quot;</span>: {
      <span class="hljs-attr">&quot;href&quot;</span>: <span class="hljs-string">&quot;https://github.com/octocat&quot;</span>,
      <span class="hljs-attr">&quot;type&quot;</span>: <span class="hljs-string">&quot;application/atom+xml&quot;</span>
    },
    <span class="hljs-attr">&quot;current_user&quot;</span>: {
      <span class="hljs-attr">&quot;href&quot;</span>: <span class="hljs-string">&quot;https://github.com/octocat.private?token=abc123&quot;</span>,
      <span class="hljs-attr">&quot;type&quot;</span>: <span class="hljs-string">&quot;application/atom+xml&quot;</span>
    },
    <span class="hljs-attr">&quot;current_user_actor&quot;</span>: {
      <span class="hljs-attr">&quot;href&quot;</span>: <span class="hljs-string">&quot;https://github.com/octocat.private.actor?token=abc123&quot;</span>,
      <span class="hljs-attr">&quot;type&quot;</span>: <span class="hljs-string">&quot;application/atom+xml&quot;</span>
    },
    <span class="hljs-attr">&quot;current_user_organization&quot;</span>: {
      <span class="hljs-attr">&quot;href&quot;</span>: <span class="hljs-string">&quot;&quot;</span>,
      <span class="hljs-attr">&quot;type&quot;</span>: <span class="hljs-string">&quot;&quot;</span>
    },
    <span class="hljs-attr">&quot;current_user_organizations&quot;</span>: [
      {
        <span class="hljs-attr">&quot;href&quot;</span>: <span class="hljs-string">&quot;https://github.com/organizations/github/octocat.private.atom?token=abc123&quot;</span>,
        <span class="hljs-attr">&quot;type&quot;</span>: <span class="hljs-string">&quot;application/atom+xml&quot;</span>
      }
    ],
    <span class="hljs-attr">&quot;security_advisories&quot;</span>: {
      <span class="hljs-attr">&quot;href&quot;</span>: <span class="hljs-string">&quot;https://github.com/security-advisories&quot;</span>,
      <span class="hljs-attr">&quot;type&quot;</span>: <span class="hljs-string">&quot;application/atom+xml&quot;</span>
    }
  }
}
</code></pre></div>
    
    
      <h4>Notes</h4>
      <ul class="mt-2">
      
        <li><a href="/en/developers/apps">Works with GitHub Apps</a></li>
      
      </ul>
    
    
  </div>
  <hr>
</div>
<h3 id="example-of-getting-an-atom-feed"><a href="#example-of-getting-an-atom-feed">Example of getting an Atom feed</a></h3>
<p>To get a feed in Atom format, you must specify the <code>application/atom+xml</code> type in the <code>Accept</code> header. For example, to get the Atom feed for GitHub security advisories:</p>