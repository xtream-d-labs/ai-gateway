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
          <div class="row">
            <div class="col s3">
              <div style="margin: 5px 0 0 2px;line-height: 3rem;">
                <span id="record-count">0</span>&nbsp;hits
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
            <li v-for="e in errors" :key="e.id" :data-id="e.id" >
              <div class="row row-header" style="padding: 13px 30px 10px 15px;"
                   :id="e.id" :data-target="'#body-'+e.id"
                   data-toggle="collapse" aria-expanded="true"
                   :aria-controls="'body-'+e.id">
                <div class="col-1">
                  <i class="material-icons" style="float: right;">chevron_right</i>
                </div>
                <div class="col-3">{{ e.occursAt }}</div>
                <div class="col-6 cut-text">{{ e.caption }}</div>
                <div class="col-2 cut-text">{{ e.owner }}</div>
              </div>
              <div :id="'body-'+e.id" class="row collapse row-body"
                   :aria-labelledby="e.id"
                   data-parent="#accordion">
                <div class="col-12">
                  <div class="row"><div class="col-12" style="margin-bottom: 20px;">
                    <h6>Error detail:</h6>
                    <span>{{ e.detail }}</a>
                  </div></div>
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
