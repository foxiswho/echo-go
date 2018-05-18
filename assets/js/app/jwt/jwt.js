define(["jquery", "jquery-form", "bootstrap"], function ($, form, bootstrap) {

    $("#foo").click(function(event) {
        jwtTester();
    });

    function jwtTester() {
        token = $('#token').val();
        url = $('#url').val();
        $.ajax({
            dataType: 'text',
            async: false,
            url: url,
            beforeSend: function(request) {
                request.setRequestHeader("Authorization", "Bearer "+token);
            },
            success: function (data) {
                // alert("成功");
            },
            error: function () {
                // alert("失败");
            },
            complete: function (xhr) {
                $('#api_data').val(xhr.responseText);
            },
        });
    }
});