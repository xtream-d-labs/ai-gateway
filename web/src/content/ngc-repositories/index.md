+++
title = "NGC Images"
css = "ngc-repositories/index.css"
js = "ngc-repositories/index.js"
+++

<main>
  <section class="container content-header">
    <div class="row">
      <div class="col s12" style="min-height: 182px;">
        <h5 class="light grey-text text-darken-2">NGC Images</h5>
        <form>
          <div class="row hide-on-small-only">
            <div class="col m6" style="padding-right: 0;">
              <div class="input-field" style="width: 90%;margin: 13px 0 -13px 0;">
                <input id="query-words" type="text" style="font-size: 1.5rem;">
                <label for="query-words">Search words</label>
              </div>
            </div>
            <div class="col m3" style="margin: 10px 0 -10px 0px;">
              <div id="categories">
                <select></select>
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
            <li v-for="image in images" :key="image.code"
                :data-ns="image.namespace" :data-nm="image.name" >
              <div class="row row-header collapsible-header" style="padding: 13px 30px 10px 15px;"
                   :id="image.name" :data-target="'#body-'+image.name"
                   data-toggle="collapse" aria-expanded="true"
                   :aria-controls="'body-'+image.name">
                <div class="col-3">
                  <i class="material-icons">cloud</i>
                  <div class="cut-text">{{ image.namespace }}</div>
                </div>
                <div class="col-9 cut-text">{{ image.name }}</div>
              </div>
              <div :id="'body-'+image.name" class="row collapse row-body"
                   :aria-labelledby="image.name"
                   data-parent="#accordion">
                <div class="col-12" style="margin-bottom: 0px;">
                  <h6 style="font-weight: 600;margin: 5px 0 10px 0;">Available versions</h6>
                </div>
                <div class="col-12" style="max-height: 195px;overflow-y: scroll;">
                  <table class="table highlight">
                    <tbody></tbody>
                  </table>
                </div>
                <div class="col-12" v-html="image.description"></div>
                <div class="clear-both"></div>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </section>

</main>
