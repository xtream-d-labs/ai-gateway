+++
title = "Notebooks"
css = "notebooks/index.css"
js = "notebooks/index.js"
+++

<main>
  <section class="container content-header">
    <div class="row">
      <div class="col s12" style="min-height: 182px;">
        <h5 class="light grey-text text-darken-2">Notebooks</h5>
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
                    <option value="1">Sort by name</option>
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
          <ul class="collapsible" id="accordion">
            <li v-for="notebook in notebooks" :key="notebook.id"
                :data-id="notebook.id" :data-url="notebook.url">
              <div class="row row-header" style="padding: 13px 30px 10px 15px;"
                   :id="notebook.id" :data-target="'#body-'+notebook.id"
                   data-toggle="collapse" aria-expanded="true"
                   :aria-controls="'body-'+notebook.id">
                <div class="col-1">
                  <i v-if="notebook.url.length > 0"
                     class="material-icons" style="float: right;">chevron_right</i>
                </div>
                <div class="col-6">
                  <div class="cut-text">{{ notebook.image }}</div>
                </div>
                <div class="col-3 cut-text">{{ notebook.started }}</div>
                <div class="col-2" style="text-align: center;">
                  <div class="label" style="margin-left: 5px;"
                      :class="notebook.classObject">{{ notebook.state }}</div>
                </div>
              </div>
              <div v-if="notebook.url.length > 0"
                   :id="'body-'+notebook.id" class="row collapse row-body"
                   :aria-labelledby="notebook.id"
                   data-parent="#accordion">
                <div class="col-12" v-if="notebook.state != 'exited'">
                  <div class="row"><div class="col-12" style="margin-bottom: 20px;">
                    <h6>Endpoint</h6>
                    <a href="#" target="_blank" class="endpoint"
                      style="font-size: 1.5rem;">{{ notebook.url }}</a>
                  </div></div>
                </div>
                <div class="col-12">
                  <div class="row"><div class="col-12" style="margin-bottom: 20px;">
                    <h6>Actions</h6>
                    <div style="margin: 20px 0 5px 0;">
                      <a v-if="trainOnCloud" class="waves-effect waves-light btn blue darken-1"
                        tabindex="0" v-on:click="train">train on cloud</a>
                      <span>&nbsp;</span>
                      <!-- <a class="waves-effect waves-light btn blue darken-1"
                        tabindex="0">Run more like this</a>
                      <span>&nbsp;</span> -->
                      <a class="waves-effect waves-light btn red lighten-2 act-stop" tabindex="0"
                        v-if="notebook.state != 'exited'" v-on:click="stop">stop</a>
                      <a class="waves-effect waves-light btn red lighten-2 act-delete"
                        tabindex="0" v-if="notebook.state == 'exited'"
                        v-on:click="del">delete</a>
                    </div>
                  </div></div>
                </div>
                <div class="col-12">
                  <div class="row"><div class="col-12" style="margin-bottom: 20px;">
                    <h6>Properties</h6>
                    <table class="table highlight">
                      <tbody>
                        <tr><td>container name:</td><td class="notebook-name"></td></tr>
                        <tr><td>started time:</td><td class="notebook-started"></td></tr>
                        <tr><td>ended time:</td><td class="notebook-ended"></td></tr>
                        <tr><td>mounted volumes:</td><td class="notebook-volumes"></td></tr>
                      </tbody>
                    </table>
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

<div id="training-dialog" class="modal popup-dialog"
    style="height: 655px;width: 60%;max-height: 100%;">
  <div class="modal-content">
    <h5>Task Definition</h5>
  </div>
  <div class="modal-footer row">
    <section class="container wait-icon" style="position: absolute;top: 40%;left: 0;display:none;">
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
    <div class="col-12" style="margin: 15px 0px 7px 0;min-height: 50px;">
      <form autocomplete="off" v-on:submit.prevent>
        <div class="form-group row">
          <label class="col-sm-3 control-label">Platform</label>
          <div class="col-sm-9 training-platform">
            <select>
              <option value="0">Kubernetes</option>
              <option value="1">Rescale</option>
            </select>
          </div>
        </div>
        <div class="form-group row considerable">
          <label class="col-sm-3 control-label">Job type</label>
          <div class="col-sm-9 training-type">
            <select>
              <option value="0">Run your iPython notebook</option>
              <option value="1">Run commmands</option>
            </select>
          </div>
        </div>
        <div class="form-group row considerable">
          <label class="col-sm-3 control-label">Notebook</label>
          <div class="col-sm-9 training-notebook">
            <select></select>
          </div>
        </div>
        <div class="form-group row">
          <label class="col-sm-3 control-label">Commands</label>
          <div class="col-sm-9">
            <input type="text" class="form-control training-cmds" v-model="cmd"
                  placeholder="python train.py --epoch 10 --gpu 1" />
          </div>
        </div>
        <div class="form-group row">
          <label class="col-sm-3 control-label">CPU</label>
          <div class="col-sm-9">
            <input type="number" class="form-control training-cpus text-right"
                   v-model="cpus" placeholder="1" />
          </div>
        </div>
        <div class="form-group row">
          <label class="col-sm-3 control-label">GPU</label>
          <div class="col-sm-9">
            <input type="number" class="form-control training-gpus text-right"
                   v-model="gpus" placeholder="0" />
          </div>
        </div>
        <div class="form-group row">
          <label class="col-sm-3 control-label">Hardware</label>
          <div class="col-sm-9 training-coretype">
            <select>
              <option value="emerald">CPU only</option>
              <option value="dolomite">GPU (Tesla V100)</option>
            </select>
          </div>
        </div>
        <div class="form-group row">
          <label class="col-sm-3 control-label">Cores</label>
          <div class="col-sm-9 training-cores">
            <select></select>
          </div>
        </div>
        <div class="clear-both"></div>
      </form>
    </div>
    <div class="row">
      <div class="col-12">
        <p style="margin-bottom: 0;">It takes 5 minutes at least to start a job on cloud.</p>
      </div>
    </div>
    <div class="clear-both"></div>
    <div class="col-12">
      <a class="waves-effect waves-light btn cancel" tabindex="0" v-on:click="close">Cancel</a>
      <a id="act-submit" class="waves-effect waves-light btn submit" tabindex="0" 
         style="float: right;color: white !important;" v-on:click="submit">Start</a>
    </div>
  </div>
</div>

<div id="notebook-modify" class="modal popup-dialog" style="height: 245px;">
  <div class="modal-content">
    <h5>Confirmation</h5>
  </div>
  <div class="modal-footer row">
    <div class="col-12" style="margin: 15px 0 20px 0;min-height: 50px;">
      <span>Is it okay to <span style="color: red;font-weight: 600;">{{ action }}</span> the following notebook?</span><br>
      <strong style="font-weight: bold;font-size: 1.5rem;">{{ name }}</strong>
    </div>
    <div class="clear-both"></div>
    <div class="col-12">
      <a class="waves-effect waves-light btn cancel" tabindex="0" v-on:click="close">Cancel</a>
      <a class="waves-effect waves-light btn blue darken-1 delete" tabindex="0"
         style="float: right;color: white !important;" v-on:click="exec">{{ action }}</a>
    </div>
  </div>
</div>
