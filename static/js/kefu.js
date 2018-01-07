var Page = {
    init: function() {
        var a = this;
        a.bindQQOnline();
    },
    bindQQOnline: function() {
        $("#floatTrigger").bind("click", function() {
            if ($("#online_qq_layer").attr("show")) {
                $("#online_qq_layer").animate({
                    right: "-140px"
                });
                $("#online_qq_layer").removeAttr("show")
            } else {
                $("#online_qq_layer").animate({
                    right: "0px"
                });
                $("#online_qq_layer").attr("show", "1")
            }
            return false
        });
        $("#online_qq_layer").animate({
            right: "-140px"
        });
        $("#online_qq_layer").removeAttr("show");
        $(document).bind("click", function(a) {
            if ($(a.target).isChildOf("#online_qq_layer") == false) {
                $("#online_qq_layer").animate({
                    right: "-140px"
                });
                $("#online_qq_layer").removeAttr("show")
            }
        });
        jQuery.fn.isChildAndSelfOf = function(a) {
            return (this.closest(a).length > 0)
        };
        jQuery.fn.isChildOf = function(a) {
            return (this.parents(a).length > 0)
        }
    }
};
$(function() {
    Page.init()
});