$(function () {
    $('#login-form').validate({
        ignore: '',
        rules: {
            phone: { required: true, number: true, rangelength: [11, 11] },
            password: { required: true }
        },
        messages: {
            phone: { required: '请填写手机号', number: '请填写正确的手机号', rangelength: '请填写11位手机号' },
            password: { required: '请填写密码' }
        },
        submitHandler: function (form) {
            var url = '/login';
            $(form).ajaxSubmit({
                url: url,
                type: 'POST',
                dataType: 'json',
                success: function (data) {
                    dialogInfo(data.message)
                    if (data.code == 1) {
                        setTimeout(function () { window.location.reload(); }, 2000);
                    } else {

                    }
                }
            });
        }
    });
});


function dialogInfo(msg) {
    var html = '';
    html += '<div class="modal fade" id="dialogInfo" tabindex="-1" role="dialog" aria-labelledby="dialogInfoTitle">';
    html += '<div class="modal-dialog" role="document">';
    html += '<div class="modal-content">';
    html += '<div class="modal-header">';
    html += '<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>';
    html += '<h4 class="modal-title" id="dialogInfoTitle">信息提示</h4>';
    html += ' </div>';
    html += '<div class="modal-body">';
    html += '<p>' + msg + '</p>';
    html += '</div>';
    //html += '<div class="modal-footer">';
    //html += ' <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>';
    //html += ' <button type="button" class="btn btn-primary">Send message</button>';
    //html += '</div>';
    html += '</div>';
    html += '</div>';
    html += '</div>';
    $('body').append(html);
    $('#dialogInfo').modal('show');
}