+++
title = "Workspaces"
css = "workspaces/index.css"
js = "workspaces/index.js"
+++

<main>
  <section class="container content-header">
    <div class="row">
      <div class="col s12" style="min-height: 182px;">
        <h5 class="light grey-text text-darken-2">Workspaces</h5>

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
            <div class="col col-md-9">
              <div class="row">
                <div class="input-field inline thin-input right col m5 s12" style="max-width: 180px;">
                  <select id="query-order-type">
                    <option value="1">Sort by name</option>
                    <option value="2" selected="selected">Sort by created time</option>
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

          <ul class="collapsible" id="accordion">
            <li v-for="workspace in workspaces" :key="workspace.path">
              <div class="row row-header" style="padding: 13px 30px 10px 15px;"
                   :id="'head-'+workspace.path" :data-target="'#'+workspace.path"
                   data-toggle="collapse" aria-expanded="true" :aria-controls="workspace.path">
                <div class="col-1">
                  <i class="material-icons" style="float: right;"
                     v-if="workspace.note.length > 0 || workspace.jobs.length > 0">folder</i>
                  <i class="material-icons" style="float: right;"
                     v-if="workspace.note.length == 0 && workspace.jobs.length == 0">folder_open</i>
                </div>
                <div class="col-7 cut-text workspace-path">{{ workspace.path }}</div>
                <div class="col-3 cut-text">{{ workspace.time }}</div>
                <div class="col-1" v-if="workspace.note.length == 0 && workspace.jobs.length == 0">
                  <a class="waves-effect waves-light btn red lighten-2" @click.stop.prevent="del">del</a>
                </div>
              </div>
              <div :id="workspace.path" class="row collapse row-body"
                   :aria-labelledby="'head-'+workspace.path"
                   data-parent="#accordion">
                <div class="col-12" style="margin-bottom: 15px;">
                  <table class="table highlight">
                    <thead>
                      <tr><th>Related notebooks</th></tr>
                    </thead>
                    <tbody v-if="workspace.note.length > 0"><tr v-for="note in workspace.note">
                      <td><a class="notebooks">{{ note }}</a></td>
                    </tr></tbody>
                    <tbody v-if="workspace.note.length == 0">
                      <tr><td>-</td></tr>
                    </tbody>
                  </table>
                </div>
                <div class="col-12" style="margin-bottom: 15px;">
                  <table class="table highlight">
                    <thead>
                      <tr><th>Related tasks</th></tr>
                    </thead>
                    <tbody v-if="workspace.jobs.length > 0"><tr v-for="job in workspace.jobs">
                      <td><a class="jobs">{{ job }}</a></td>
                    </tr></tbody>
                    <tbody v-if="workspace.jobs.length == 0">
                      <tr><td>-</td></tr>
                    </tbody>
                  </table>
                </div>
                <div class="col-12" style="margin-bottom: 15px;">
                  <table class="table highlight">
                    <thead>
                      <tr><th>Absolute path</th></tr>
                    </thead>
                  </table>
                  <p style="margin: 10px 0 0 7px;max-width: 95%;white-space: pre-wrap;"
                    >{{ workspace.full }}</p>
                </div>
                <div class="col-12" style="margin-bottom: 0;">
                  <table class="table highlight">
                    <thead>
                      <tr><th>Created</th></tr>
                    </thead>
                  </table>
                  <p style="margin: 10px 0 0 7px;">{{ workspace.time }}</p>
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

<div id="workspace-delete" class="modal popup-dialog" style="height: 245px;">
  <div class="modal-content">
    <h5>Confirmation</h5>
  </div>
  <div class="modal-footer row">
    <div class="col-12" style="margin: 15px 0 22px 0;">
      <span>Is it okay to remove the following workspace?</span><br>
      <strong style="font-weight: bold;font-size: 1.5rem;"></strong>
    </div>
    <div class="clear-both"></div>
    <div class="col-12">
      <a class="waves-effect waves-light btn cancel" tabindex="0">Cancel</a>
      <a class="waves-effect waves-light btn blue darken-1 delete" tabindex="0"
         style="float: right;color: white !important;">OK</a>
    </div>
  </div>
</div>
