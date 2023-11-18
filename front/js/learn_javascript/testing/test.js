'use strict'

describe("pow", function () {


    describe("�������� x � ������� 3", function () {

        function makeTest(x) {
            let expected = x * x * x;
            it(`${x} � ������� 3 ����� ${expected}`, function () {
                assert.equal(pow(x, 3), expected);
            });
        }

        for (let x = 1; x <= 5; x++) {
            makeTest(x);
        }

    });

    it("��� ������������� n ���������� NaN", function () {
        assert.isNaN(pow(2, -1));
    });

    it("��� ������� n ���������� NaN", function () {
        assert.isNaN(pow(2, 1.5));
    });

});