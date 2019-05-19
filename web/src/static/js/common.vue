// <script>
window.app = {
  stop: function (e) {
    e = e || window.event;
    if (!e)
      return false;
    e.cancelBubble = true;
    if (e.stopPropagation)
      e.stopPropagation();
    e.returnValue = false;
    if (e.preventDefault)
      e.preventDefault();
    return e;
  },
  shouldExit: function (res, err) {
    if (! $.isEmptyObject(err)) {
      return true;
    }
    if (res && res.body && res.body.code == 500 && res.body.message == 'invalid token') {
      return true;
    }
    return false;
  },
  trim: function (value) {
    return value ? app.singleSpace(value).replace(/\s/g, '') : '';
  },
  singleSpace: function (value) {
    return value.replace(/　/g, ' ').replace(/\s+/g, ' '); // eslint-disable-line no-irregular-whitespace
  },
  match: function (values, searchWord) {
    var matched = true;
    if (values == undefined) values = '';
    var searchWords = app.singleSpace(searchWord).split(' ');

    if ($.isArray(values)) {
      $.map(searchWords, function (word) {
        if (word == '') return;
        var nested = false;
        $.map(values, function (value) {
          nested |= ((value+'').indexOf(word) > -1);
        });
        matched &= nested;
      });
    } else {
      $.map(searchWords, function (word) {
        if (word == '') return;
        matched &= ((values+'').indexOf(word) > -1);
      });
    }
    return matched;
  },
  comma: function (value, unit) {
    if (value === undefined) return "";
    if (unit == 'yen') {
      return value.toString().replace(/(\d)(?=(\d{3})+$)/g, '$1,') + '円';
    } else if (unit == 'byte') {
      if      (value >= 1073741824) {return (value/1073741824).toFixed(2)+' GB';}
      else if (value >= 1048576)    {return (value/1048576).toFixed(2)+' MB';}
      else if (value >= 1024)       {return (value/1024).toFixed(2)+' KB';}
      else if (value >  1)          {return value+' bytes';}
      else if (value == 1)          {return value+' byte';}
      else                          {return '0 byte';}
    }
    return value.toString().replace(/(\d)(?=(\d{3})+$)/g, '$1,');
  },
  query: function (key, def) {
    key = key.replace(/[\[]/, "\\[").replace(/[\]]/, "\\]"); // eslint-disable-line no-useless-escape
    def = def ? def : "";
    var regex = new RegExp("[\\?&]" + key + "=([^&#]*)"),
        results = regex.exec(location.search);
    return results === null ? def : decodeURIComponent(results[1].replace(/\+/g, " "));
  },
  date: {
    format: function (d, lang) {
      if (lang == 'en') {
        // Mon, 15 Jun 2009 20:45:30
        var weekday = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
        var months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
        return weekday[d.getDay()]+', '+_fill(d.getDate())+' '+months[d.getMonth()]+' '+
            _fill(d.getYear() + 1900)+' '+_fill(d.getHours())+
            ":"+_fill(d.getMinutes())+":"+_fill(d.getSeconds());
      }
      // 2009/01/15 20:45:30
      return  _fill(d.getYear() + 1900)+"/"+_fill(d.getMonth() + 1)+
          "/"+_fill(d.getDate())+' '+_fill(d.getHours())+
          ":"+_fill(d.getMinutes())+":"+_fill(d.getSeconds());
      function _fill(value) {
        return (value < 10) ? '0'+value : value;
      }
    },
    fromto: function (start, end, short, lang) {
      var s, e;
      if (lang == 'en') {
        s = _datetimeEn(start, short);
        e = end ? _datetimeEn(end, short) : '';
        if (s == e) {
          e = '';
        } else {
          if (short) {
            e = e.replace(s.substring(0, 6), '');
            return s + ' ~ ' + e;
          }
          e = e.replace(s.substring(0, 16), '');
        }
        return (e == '') ? s + ' ~ ' : s + ' ~ ' + e;
      }
      s = _datetimeJa(start, short);
      e = end ? _datetimeJa(end, short) : '';
      if (s == e) {
        e = '';
      } else {
        if (short) {
          e = e.replace(s.substring(0, 6), '');
          return s + ' ~ ' + e;
        }
        e = e.replace(s.substring(0, 11), '');
      }
      if (e == '') {
        return s + ' ~ ';
      }
      return (s + ' ~ ' + e);

      function _datetimeEn(date, short) {
        date = app.date.format(new Date(date), 'en');
        return short ? date.substring(5, 11) + date.substring(16, 22) : date.substring(0, 22);
      }
      function _datetimeJa(date, short) {
        date = app.date.format(new Date(date));
        return short ? date.substring(5, 16) : date.substring(0, 16);
      }
    },
    lastMidnight: function () {
      return new Date(new Date().setHours(0,0,0,0));
    }
  },
  textarea: function (opt){
    var textarea, height;
    if (opt.id) textarea = document.getElementById(opt.id);
    if (opt.textarea) textarea = opt.textarea;
    if (! textarea) return;

    opt.min = opt.min ? opt.min : 30;
    height = parseInt(textarea.style.height, 10);
    textarea.style.height = ((height < opt.min) ? opt.min : height) + "px";

    height = parseInt(textarea.scrollHeight, 10);
    if (opt.max && (opt.max < height)) height = opt.max;
    textarea.style.height = height + "px";
  },
  ui: {
    datetime: function (arg) {
      var input = $('#'+arg.id).attr('data-value', app.date.format(arg.date)).pickadate({
        monthsFull: ['1 月', '2 月', '3 月', '4 月', '5 月', '6 月', '7 月', '8 月', '9 月', '10 月', '11 月', '12 月'],
        weekdaysShort: ['日', '月', '火', '水', '木', '金', '土'],
        weekdaysLetter: ['日', '月', '火', '水', '木', '金', '土'],
        format: 'yyyy年m月d日(ddd)',
        formatSubmit: 'yyyy-mm-dd',
        selectMonths: true,
        selectYears: 3,
        max: arg.max !== undefined ? arg.max : false,
        min: arg.min !== undefined ? arg.min : false,
        onOpen: function () {
          app.ui.relocate({triggerId: '#'+arg.id, followerId: '#'+arg.id+'_root .picker__frame'});
          $('#'+arg.id+'_root').animate({opacity: 1});
        },
        onSet: function (value) {
          if (value && value.select) {
            arg.onSet && arg.onSet(value.select);
            this.close();
          }
        },
        onClose: function () {
          $(document.activeElement).blur();
          $('#'+arg.id+'_root').css({opacity: 0});
          arg.onClose && arg.onClose();
        }
      });
      return input.pickadate('picker');
    },
    relocate: function (arg) {
      var input = $(arg.triggerId);
      var frame = $(arg.followerId);

      var top = input.offset().top + input.height() + 4 - $(window).scrollTop();
      if (arg.right) {
        frame.css({top: top + 'px', left: (input.offset().left + input.width() - 1
          + parseInt(input.css('paddingLeft'), 10) - frame.width()) + 'px'});
      } else {
        frame.css({top: top + 'px', left: (input.offset().left + 2) + 'px'});
      }
    }
  },
  storage: {
    get: function(key){
      var value = localStorage.getItem(key);
      if (!value){
        return null;
      }
      return JSON.parse(value);
    },
    set: function(key, value){
      var json = JSON.stringify(value);
      localStorage.setItem(key, json);
    },
    remove: function(key){
      localStorage.removeItem(key);
    }
  },
  refreshMenus: function (config) {
    if (config.useNGC || config.usePrivateRegistry) {
      $('#menu-repositories').fadeIn();
    } else {
      $('#menu-repositories').hide();
    }
    if (config.useNGC) {
      $('#menu-ngc-repo').fadeIn();
    } else {
      $('#menu-ngc-repo').hide();
    }
    if (config.usePrivateRegistry) {
      $('#menu-prv-repo').fadeIn();
    } else {
      $('#menu-prv-repo').hide();
    }
    if (config.useRescale || config.useKubernetes) {
      $('#menu-tasks').fadeIn();
    } else {
      $('#menu-tasks').hide();
    }
  }
};

$('#signout').click(function (e) {
  app.stop(e);
  auth.exit();
});

$(document).ready(function () {
  var session = auth.session();
  var c = config.get();
  // if ($('body').hasClass('skin-black') && c.mustSignedIn) {
  //   if (! session.signedIn) {
  //     window.location.href = '/';
  //     return;
  //   }
  // }
  app.refreshMenus(c);
  if (session.username != undefined && session.username != '---') {
    $('.user-name').text(session.username);
    $('#signout').parent().show();
  }
  $('select').formSelect();
  var modals = $('.modal').detach();
  $('body').append(modals);
  modals.modal();
});
// </script>
