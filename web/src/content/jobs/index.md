+++
title = "Training Tasks"
css = "jobs/index.css"
js = "jobs/index.js"
+++

<main>
  <section class="container content-header">
    <div class="row">
      <div class="col s12" style="min-height: 182px;">
        <h5 class="light grey-text text-darken-2">Training Tasks</h5>

        <form>
          <div class="row hide-on-small-only">
            <div class="col m12" style="padding-right: 0;">
              <div class="input-field" style="width: 90%;margin: 13px 0 -13px 0;">
                <input id="query-words" type="text" style="font-size: 1.5rem;">
                <label for="query-words">Search words</label>
              </div>
            </div>
          </div>
          <div class="clear-both"></div>

          <div class="row">
            <div class="col s3">
              <div style="margin: 5px 0 0 2px;line-height: 3rem;">
                <span id="record-count">0</span>&nbsp;hits
              </div>
            </div>
            <div class="col s9">
              <div class="row">
                <div class="input-field inline thin-input right col m5 s12" style="max-width: 180px;">
                  <select id="query-order-type">
                    <option value="1">Sort by status</option>
                    <option value="2" selected="selected">Sort by started time</option>
                  </select>
                </div>
              </div>
            </div>
          </div>
        </form>
      </div>
    </div>
  </section>

  <section class="container main">
    <div class="row">
      <div class="col s12" style="margin-bottom: 15px;">
        <div id="data">

          <ul class="collapsible" data-collapsible="accordion">
            <li v-for="job in jobs" :key="job.id" :data-id="job.id" >
              <div class="collapsible-header row">
                <div class="col-1">
                  <i class="material-icons" style="float: right;"
                     v-if="job.status == 'preparing'">cached</i>
                  <i class="material-icons" style="float: right;"
                     v-if="job.status == 'sending'">cloud_upload</i>
                  <i class="material-icons" style="float: right;"
                     v-if="job.status == 'running'">fast_forward</i>
                  <i class="material-icons" style="float: right;"
                     v-if="job.status == 'done'">done_outline</i>
                </div>
                <div class="col-6 cut-text">{{ job.image }}</div>
                <div class="col-3 cut-text">{{ job.started }}</div>
                <div class="col-2" style="text-align: center;">
                  <div class="label" style="margin-left: 5px;"
                      :class="job.classObject">{{ job.status }}</div>
                </div>
              </div>
              <div class="collapsible-body row">
                <div class="col-12" style="margin-bottom: 15px;">
                  <h6>Status</h6>
                  <p style="font-size: 1.5rem;margin: 0 0 5px 6px;">{{ job.statusMore }}</p>
                </div>
                <div class="col-12" style="margin-bottom: 20px;">
                  <h6>Image</h6>
                  <p style="margin: 10px 0 0 7px;" v-html="job.imageHref"></p>
                </div>
                <div class="col-12" style="margin-bottom: 20px;">
                  <h6>Commands</h6>
                  <p style="font-size: 1.5rem;margin: 10px 0 0 7px;">{{ job.commands }}</p>
                </div>
                <div class="col-12" style="margin-bottom: 20px;">
                  <h6>Actions</h6>
                  <div style="margin: 20px 0 5px 7px;">
                    <span v-if="job.status == 'preparing'">-</span>
                    <a class="waves-effect waves-light btn red lighten-2 act-stop" tabindex="0"
                       v-if="job.status == 'running'">stop</a>
                    <a class="waves-effect waves-light btn red lighten-2 act-delete" tabindex="0"
                       v-if="job.status == 'done'" v-on:click="del">delete</a>
                  </div>
                </div>
                <div class="col-12" style="margin-bottom: 0px;">
                  <h6>Properties</h6>
                  <table class="table highlight">
                    <tbody>
                      <tr><td>task id:</td><td>{{ job.id }}</td></tr>
                      <tr><td>mounted volumes:</td><td v-html="job.mounts"></td></tr>
                      <tr><td>started time:</td><td>{{ job.started }}</td></tr>
                    </tbody>
                  </table>
                </div>
                <div class="clear-both"></div>
              </div>
            </li>
          </ul>

        </div>
      </div>
    </div>

  </section>
</main>

<div id="job-modify" class="modal popup-dialog" style="height: 245px;">
  <div class="modal-content">
    <h5>Confirmation</h5>
  </div>
  <div class="modal-footer row">
    <div class="col-12" style="margin: 15px 0 20px 0;">
      <span>Is it okay to <span style="color:red;">{{ action }}</span> the following job?</span><br>
      <strong style="font-weight: bold;font-size: 1.5rem;">{{ id }}</strong>
    </div>
    <div class="col-12">
      <a class="waves-effect waves-light btn cancel" tabindex="0" v-on:click="close">Cancel</a>
      <a class="waves-effect waves-light btn blue darken-1 delete" tabindex="0"
         style="color: white !important;" v-on:click="exec">OK</a>
    </div>
  </div>
</div>
