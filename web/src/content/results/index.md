+++
title = "Task Result"
css = "tasks/result.css"
js = "tasks/result.js"
+++

<main>
  <section class="container content-header">
    <div class="row">
      <div class="col s12" style="margin-bottom: 15px;">
        <form id="form-job-result" style="min-height: 75px;">
          <h5 class="form-signin-heading" style="margin-top: 0;">Task Result</h5>
          <hr/>
          <table class="table highlight">
            <tbody>
              <tr><td>Status:</td><td>{{ task.statusMore }}</td></tr>
              <tr><td>Base Image:</td><td v-html="task.imageHref"></td></tr>
              <tr><td>Copied volumes:</td><td v-html="task.mounts"></td></tr>
              <tr><td>Commands:</td><td>{{ task.commands }}</td></tr>
              <tr><td>Platform:</td><td>{{ task.platform }}</td></tr>
              <tr v-if="task.link != ''"><td></td><td>
                <a v-bind:href='task.link' target="_blank">{{ task.link }}</a>
              </td></tr>
              <tr><td>Started:</td><td>{{ task.started }}</td></tr>
              <tr><td>Ended:</td><td>{{ task.ended }}</td></tr>
            </tbody>
          </table>
          <div class="clear-both"></div>
        </form>
      </div>
    </div>
    <div id="logs" class="row">
      <div class="col s12" style="margin-bottom: 15px;padding-bottom: 20px;">
        <table class="table highlight">
          <thead>
            <tr><th>Task Log</th></tr>
          </thead>
          <tbody>
            <tr v-for="log in logs" :key="log.seq">
              <td><div style="min-width: 130px;">{{ log.time }}</div></td>
              <td>{{ log.log }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div id="files" class="row">
      <div class="col s12" style="margin-bottom: 15px;padding-bottom: 20px;">
        <table class="table highlight">
          <thead>
            <tr><th>Output files</th></tr>
          </thead>
          <tbody>
            <tr v-for="file in files" :key="file.seq" :data-url="file.url" >
              <td><a href="#" @click.stop.prevent="download">{{ file.name }}</a></td>
              <td>{{ file.size }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      </div>
    </div>
  </section>

  <section class="container wait-icon">
    <div style="margin: 30px auto 0 auto;width: 75px;">
      <div class="preloader-wrapper big active">
        <div class="spinner-layer spinner-green-only">
          <div class="circle-clipper left">
            <div class="circle"></div>
          </div>
          <div class="gap-patch">
            <div class="circle"></div>
          </div>
          <div class="circle-clipper right">
            <div class="circle"></div>
          </div>
        </div>
      </div>
    </div>
  </section>
</main>
