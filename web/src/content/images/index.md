+++
title = "Local Images"
css = "images/index.css"
js = "images/index.js"
+++

<main>
  <section class="container content-header">
    <div class="row">
      <div class="col s12" style="min-height: 182px;">
        <h5 class="light grey-text text-darken-2">Local Images</h5>

        <div style="position: absolute;top: 60px;right: 45px;z-index: 1001;">
          <div style="margin: -5px 7px 0 0;text-align: right;">
            <a id="act-pull" href="#">Download a new image</a>
          </div>
        </div>
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
                    <option value="1" selected="selected">Sort by name</option>
                    <option value="2">Sort by created time</option>
                    <option value="3">Sort by size</option>
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
            <li v-for="image in images" :key="image.id" :data-id="image.id" >
              <div class="collapsible-header row" style="padding: 13px 0 10px 0;">
                <div class="col-5 image-tags">
                  <i class="material-icons">layers</i>
                  <div class="cut-text image-tag">{{ image.tag }}</div>
                </div>
                <div class="col-2 text-right">{{ image.size }}</div>
                <div class="col-3 text-center">{{ image.created }}</div>
                <div class="col-2" v-if="image.size != ''">
                  <a class="waves-effect waves-light btn blue darken-1" @click.stop.prevent="run">run</a>
                  <a class="waves-effect waves-light btn red lighten-2" @click.stop.prevent="del">del</a>
                </div>
              </div>
            </li>
          </ul>

        </div>
      </div>
    </div>

  </section>
</main>

<div id="image-dialog" class="modal popup-dialog" style="height: 270px;">
  <div class="modal-content">
    <h5>Image name</h5>
  </div>
  <div class="modal-footer row">
    <div class="col-12" style="height: 111px;">
      <form class="input-field" autocomplete="off" v-on:submit.prevent>
        <p>You can use any docker images which contains Python & pip.</p>
        <input type="text" class="form-control" required style="font-size: 1.7rem;"
               placeholder="tensorflow/tensorflow:1.13.1-py3"
               v-model="name" v-on:input="$v.name.$touch" v-on:change="nameChanged"
               v-bind:class="{invalid: $v.name.$error, valid: $v.name.$dirty && !$v.name.$invalid}" />
        <p class="errors"><transition name="fade">
          <span v-if="msgName != ''">{{ msgName }}</span>
        </transition></p>
      </form>
    </div>
    <div class="clear-both"></div>
    <a class="waves-effect waves-light btn submit" tabindex="0" v-on:click="submit">Download</a>
    <a class="waves-effect waves-light btn cancel" tabindex="0" v-on:click="close">Cancel</a>
  </div>
</div>

<div id="image-delete" class="modal popup-dialog" style="height: 270px;">
  <div class="modal-content">
    <h5>Confirmation</h5>
  </div>
  <div class="modal-footer row">
    <div class="col-12" style="margin: 15px 0 22px 0;min-height: 78px;">
      <span>Is it okay to remove the following image?</span><br>
      <strong style="font-weight: bold;font-size: 1.5rem;"></strong>
    </div>
    <div class="clear-both"></div>
    <div class="col-12">
      <a class="waves-effect waves-light btn cancel" tabindex="0">Cancel</a>
      <a class="waves-effect waves-light btn blue darken-1 delete" tabindex="0"
         style="float: right;color: white !important;">OK</a>
    </div>
  </div>
  <div class="clear-both"></div>
</div>

<div id="run-dialog" class="modal popup-dialog"
    style="height: 320px;width: 60%;max-height: 85%;">
  <div class="modal-content">
    <h5>Run configurations</h5>
  </div>
  <div class="modal-footer row" style="margin: 0;">
    <div class="col-12" style="margin-top: 13px;">
      <form autocomplete="off" style="min-height: 147px;" v-on:submit.prevent>
        <div class="form-group row">
          <label class="col-sm-2 control-label">Workspace</label>
          <div class="col-sm-10 workspace-type">
            <select>
              <option value="0">Create its own workspace</option>
              <option value="1">Use an existing workspace</option>
            </select>
          </div>
        </div>
        <div class="form-group row considerable">
          <label class="col-sm-2 control-label">&nbsp;</label>
          <div class="col-sm-10 workspaces">
            <select></select>
          </div>
        </div>
        <div class="clear-both"></div>
      </form>
    </div>
    <div class="clear-both"></div>
    <div class="col-12">
      <a class="waves-effect waves-light btn cancel" tabindex="0" v-on:click="close">Cancel</a>
      <a class="waves-effect waves-light btn submit" tabindex="0" v-on:click="submit"
         style="float: right;color: white !important;">Run</a>
    </div>
  </div>
  <div class="clear-both"></div>
</div>
