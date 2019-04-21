+++
title = "Remote Images"
css = "repositories/index.css"
js = "repositories/index.js"
+++

<main>
  <section class="container content-header">
    <div class="row">
      <div class="col s12" style="min-height: 182px;">
        <h5 class="light grey-text text-darken-2">Remote Images</h5>
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
                <div class="col-1">
                  <i class="material-icons">cloud</i>
                </div>
                <div class="col-11 cut-text">{{ image.name }}</div>
              </div>
              <div :id="'body-'+image.name" class="row collapse row-body"
                   :aria-labelledby="image.name" data-parent="#accordion">
                <div class="col-12" style="margin-bottom: 0px;">
                  <h6 style="font-weight: 600;margin: 5px 0 10px 0;">Available versions</h6>
                </div>
                <div class="col-12" style="max-height: 195px;overflow-y: scroll;">
                  <div class="progress">
                    <div class="indeterminate"></div>
                  </div>
                  <table class="table highlight">
                    <tbody></tbody>
                  </table>
                </div>
                <div class="col-12" style="margin-bottom: 0px;">
                  <h6 style="font-weight: 600;margin: 25px 0 10px 0;">Description</h6>
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
