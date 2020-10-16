+++
title = "About AI Gateway"
css = "about/index.css"
js = "about/index.js"
+++
<main>
  <section class="container content-header">
    <div class="row">
      <div class="col s12" style="min-height: 182px;padding-bottom: 25px;margin-bottom: 40px;">
        <h4 class="form-signin-heading">AI Gateway</h4>
        <hr/>
        <h6 style="margin: 30px 0 7px 3px;font-weight: 600;">Repository</h6>
        <ul>
          <li>
            <div class="row" style="padding: 5px 0;">
              <div class="col-3"></div>
              <div class="col-9">
                <a href="https://github.com/xtream-d-labs/ai-gateway"
                   target="_blank">github.com/xtream-d-labs/ai-gateway</a>
              </div>
            </div>
          </li>
        </ul>
        <h6 style="margin: 20px 0 12px 3px;font-weight: 600;">Versions</h6>
        <ul id="data">
          <li>
            <div class="row" style="padding: 5px 0;">
              <div class="col-3">Current:</div>
              <div class="col-9">{{ current.version }}</div>
            </div>
          </li>
          <li>
            <div class="row" style="padding: 5px 0;">
              <div class="col-3"></div>
              <div class="col-9">{{ current.date }}</div>
            </div>
          </li>
          <li>
            <div class="row" style="padding: 5px 0;">
              <div class="col-3">Latest:</div>
              <div class="col-9">{{ latest.version }}</div>
            </div>
          </li>
          <li>
            <div class="row" style="padding: 5px 0;">
              <div class="col-3"></div>
              <div class="col-9">{{ latest.date }}</div>
            </div>
          </li>
          <li>
            <div class="row" style="padding: 5px 0;">
              <div class="col-3"></div>
              <div class="col-9">
                <a href="https://s3-ap-northeast-1.amazonaws.com/ai-gateway/docker-compose.yml">docker-compose.yml with port 80</a>
              </div>
            </div>
          </li>
          <li>
            <div class="row" style="padding: 5px 0;">
              <div class="col-3"></div>
              <div class="col-9">
                <a href="https://s3-ap-northeast-1.amazonaws.com/ai-gateway/docker-compose-8080.yml">docker-compose.yml with port 8080</a>
              </div>
            </div>
          </li>
          <li>
            <div class="row" style="padding: 5px 0;">
              <div class="col-3">Archived:</div>
              <div class="col-9">
                <a href="https://hub.docker.com/r/aigateway/api/tags" target="_blank">images</a>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </section>
</main>
