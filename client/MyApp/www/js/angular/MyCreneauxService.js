app.service('MyCreneaux', function($q,MyDatas, Fds, Storage) {

    console.log('Creating MyCreneaux service...');

    var deferred = $q.defer();

    var factory = function() {
        var that = this;

        this.myCreneaux = [];

        this.GetCreneaux = function(callback) {
            this.syncCreneaux();
            callback(this.myCreneaux);
        };

        this.syncCreneaux = function() {
            console.log("MyCreneaux:: syncCreneaux() ...");
            var mesUes = MyDatas.GetUes();
            Fds.GetCreneaux(mesUes, function(out) {
                that.myCreneaux = out;
                Storage.updateDb('mycreneaux', that.myCreneaux);
                console.log("MyCreneaux:: syncCreneaux() done");
            });
        };

        this.loadCreanauxFromLocalStorage = function() {
            Storage.get('mycreneaux').then(function(doc) {
                that.myCreneaux = doc.list;
                console.log("hey:");
                console.log(that.myCreneaux);
            }).then(function(response) {
                console.log("MyCreneaux:: loadCreanauxFromLocalStorage done");
                deferred.resolve();
            }).catch(function (err) {
                console.log("MyCreneaux:: loadCreanauxFromLocalStorage error");
                console.log(err);
                deferred.resolve();
            });
        };


    };

    var o = new factory();
    o.loadCreanauxFromLocalStorage();
    o.promise = deferred.promise;
    return o;
});
