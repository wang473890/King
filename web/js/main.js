(function ($) {
    "use strict";


    /*==================================================================
    [ Validate ]*/
    var input = $('.validate-input .input100');

    $('.validate-form').on('submit', function () {

        var check = true;

        for (var i = 0; i < input.length; i++) {
            if (validate(input[i]) == false) {
                showValidate(input[i]);
                check = false;
            }
        }
        if(!check){
            return false;
        }
        $.ajax({
            type: 'POST',
            url: "localhost:8000/login",
            data: {name:$('[name=name]').val().trim(),pass:$('[name=pass]').val().trim()},
            success: function(){
                alert('success');
            },
            dataType: 'json'
        });
        alert($('[name=name]').val().trim())

        // $.post(
        //     url: "localhost:8000/login",
        //     data: {name:$('[name=name]').val().trim(),pass:$('[name=pass]').val().trim()},
        //     function(result){
        // alert("登陆成功")
        // });

    });


    $('.validate-form .input100').each(function () {
        $(this).focus(function () {
            hideValidate(this);
        });
    });

    function validate(input) {
        if ($(input).attr('type') == 'name' || $(input).attr('name') == 'name') {
            // if($(input).val().trim().match(/^([a-zA-Z0-9_\-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([a-zA-Z0-9\-]+\.)+))([a-zA-Z]{1,5}|[0-9]{1,3})(\]?)$/) == null) {
            if (6 > $(input).val().trim().length || 12 < $(input).val().trim().length) {
                return false;
            }
        }
        else {
            if ($(input).val().trim() == '') {
                return false;
            }
        }
    }

    function showValidate(input) {
        var thisAlert = $(input).parent();

        $(thisAlert).addClass('alert-validate');
    }

    function hideValidate(input) {
        var thisAlert = $(input).parent();

        $(thisAlert).removeClass('alert-validate');
    }


})(jQuery);