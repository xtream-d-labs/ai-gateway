+++
title = "Errors"
css = "system/errors.css"
js = "system/errors.js"
+++

<main>
  <section class="container content-header">
    <div class="row">
      <div class="col s12" style="min-height: 182px;">
        <h5 class="light grey-text text-darken-2">Errors</h5>
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
        </form>
      </div>
    </div>
  </section>

  <section class="container main">
    <div class="row">
      <div class="col s12" style="margin-bottom: 15px;">
        <div id="data">
          <ul class="collapsible" id="accordion">
            <li v-for="job in jobs" :key="job.id" :data-id="job.id" >
              <div class="row row-header" style="padding: 13px 30px 10px 15px;"
                   :id="job.id" :data-target="'#body-'+job.id"
                   data-toggle="collapse" aria-expanded="true"
                   :aria-controls="'body-'+job.id">
                <div class="col-1">
                  <i class="material-icons" style="float: right;width: 24px;"
                     v-if="job.status == 'preparing'">cached</i>
                  <i class="material-icons" style="float: right;width: 24px;"
                     v-if="job.status == 'sending'">cloud_upload</i>
                  <i class="material-icons" style="float: right;width: 24px;"
                     v-if="job.status == 'running'">fast_forward</i>
                  <i class="material-icons" style="float: right;width: 24px;"
                     v-if="job.status == 'done'">done_outline</i>
                </div>
                <div class="col-6 cut-text">{{ job.image }}</div>
                <div class="col-3 cut-text">{{ job.started }}</div>
                <div class="col-2" style="text-align: center;">
                  <div class="label" style="margin-left: 5px;"
                      :class="job.classObject">{{ job.status }}</div>
                </div>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </section>
</main>
