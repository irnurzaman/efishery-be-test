import pytest
import json
from service import Service

@pytest.fixture(scope='module')
def service():
    return Service('API_KEY', 'RESOURCE_ADDRES', 10)

class TestFiltering:
    def test_all_invalid(self, service: Service):
        data = [{'uuid':None, 'komoditas':None, 'area_provinsi':None, 'area_kota':None, 'size':None, 'price':None, 'tgl_parsed':None, 'timestamp':None},
                {'uuid':'57d8839f-6583-46c9-8e5d-67175363332d', 'komoditas':'Teri', 'area_provinsi':'BANTEN', 'area_kota':'CIREBON', 'size':'50', 'price':None, 'tgl_parsed':'2020-06-09T15:08:58.122Z', 'timestamp':'1591690138122'},
                {'uuid':None, 'komoditas':None, 'area_provinsi':None, 'area_kota':None, 'size':None, 'price':None, 'tgl_parsed':None, 'timestamp':None}]
        result = service.filtering(data)
        assert len(result) == 0

    def test_all_valid(self, service: Service):
        data = [{'uuid':'a9e10e7d-d385-49b5-ac55-fe6d04f300af', 'komoditas':'Patin Black', 'area_provinsi':'JAWA TENGAH', 'area_kota':'PURWOREJO', 'size':'120', 'price':'20000', 'tgl_parsed':'2021-07-25T04:33:06.940Z', 'timestamp':'1627187586940'},
                {'uuid':'ff8e461c-910b-4d11-b531-8494c1c7da2e', 'komoditas':'Gabus', 'area_provinsi':'JAWA BARAT', 'area_kota':'JAWA BARAT', 'size':'70', 'price':'30000', 'tgl_parsed':'2020-06-28T14:53:22+0800', 'timestamp':'1593327202'},
                {'uuid':'8616b242-a45d-4718-873c-3cb43bdc0f08', 'komoditas':'Cumi', 'area_provinsi':'JAWA TIMUR', 'area_kota':'PROBOLINGGO', 'size':'50', 'price':'45000', 'tgl_parsed':'2020-07-27T01:11:05+07:00', 'timestamp':'1595787065171'}]
        result = service.filtering(data)
        assert len(result) == 3

    def test_partial_valid(self, service: Service):
        data = [{'uuid':'a9e10e7d-d385-49b5-ac55-fe6d04f300af', 'komoditas':'Patin Black', 'area_provinsi':'JAWA TENGAH', 'area_kota':'PURWOREJO', 'size':'120', 'price':'20000', 'tgl_parsed':'2021-07-25T04:33:06.940Z', 'timestamp':'1627187586940'},
                {'uuid':'ff8e461c-910b-4d11-b531-8494c1c7da2e', 'komoditas':'Gabus', 'area_provinsi':'JAWA BARAT', 'area_kota':'JAWA BARAT', 'size':'70', 'price':'30000', 'tgl_parsed':'2020-06-28T14:53:22+0800', 'timestamp':'1593327202'},
                {'uuid':'57d8839f-6583-46c9-8e5d-67175363332d', 'komoditas':'Teri', 'area_provinsi':'BANTEN', 'area_kota':'CIREBON', 'size':'50', 'price':None, 'tgl_parsed':'2020-06-09T15:08:58.122Z', 'timestamp':'1591690138122'}]
        result = service.filtering(data)
        assert len(result) == 2

class TestAggregating:
    def test_aggregate_fields(self, service: Service):
        errors = []
        data = [{'uuid':'a9e10e7d-d385-49b5-ac55-fe6d04f300af', 'komoditas':'Patin Black', 'area_provinsi':'JAWA TENGAH', 'area_kota':'PURWOREJO', 'size':'120', 'price':'20000', 'tgl_parsed':'2021-07-26T04:33:06.940Z', 'timestamp':'1627187586940'},
                {'uuid':'ff8e461c-910b-4d11-b531-8494c1c7da2e', 'komoditas':'Gabus', 'area_provinsi':'JAWA BARAT', 'area_kota':'JAWA BARAT', 'size':'70', 'price':'30000', 'tgl_parsed':'2020-06-28T14:53:22+0800', 'timestamp':'1593327202'},
                {'uuid':'8616b242-a45d-4718-873c-3cb43bdc0f08', 'komoditas':'Cumi', 'area_provinsi':'JAWA TIMUR', 'area_kota':'PROBOLINGGO', 'size':'50', 'price':'45000', 'tgl_parsed':'2020-07-27T01:11:05+07:00', 'timestamp':'1595787065171'}]
        post_filter = service.filtering(data)
        result = service.aggregating(post_filter)

        if 'end_week' not in result['result'][0]:
            errors.append('Missing start_week field')
        if 'area_provinsi' not in result['result'][0]:
            errors.append('Missing area_provinsi field')
        if 'min_size' not in result['result'][0]:
            errors.append('Missing min_size field')
        if 'max_size' not in result['result'][0]:
            errors.append('Missing max_size field')
        if 'med_size' not in result['result'][0]:
            errors.append('Missing med_size field')
        if 'mean_size' not in result['result'][0]:
            errors.append('Missing mean_size field')
        if 'min_price' not in result['result'][0]:
            errors.append('Missing min_price field')
        if 'max_price' not in result['result'][0]:
            errors.append('Missing max_price field')
        if 'med_price' not in result['result'][0]:
            errors.append('Missing med_price field')
        if 'mean_price' not in result['result'][0]:
            errors.append('Missing start_week field')

        assert not errors, ','.join(errors)

    def test_aggregate_values(self, service: Service):
        errors = []
        data = [{'uuid':'a9e10e7d-d385-49b5-ac55-fe6d04f300af', 'komoditas':'Patin Black', 'area_provinsi':'JAWA TIMUR', 'area_kota':'PURWOREJO', 'size':'120', 'price':'20000', 'tgl_parsed':'2021-07-26T04:33:06.940Z', 'timestamp':'1627187586940'},
                {'uuid':'ff8e461c-910b-4d11-b531-8494c1c7da2e', 'komoditas':'Gabus', 'area_provinsi':'JAWA BARAT', 'area_kota':'JAWA BARAT', 'size':'70', 'price':'30000', 'tgl_parsed':'2021-06-28T14:53:22+0800', 'timestamp':'1593327202'},
                {'uuid':'8616b242-a45d-4718-873c-3cb43bdc0f08', 'komoditas':'Cumi', 'area_provinsi':'JAWA TIMUR', 'area_kota':'PROBOLINGGO', 'size':'50', 'price':'45000', 'tgl_parsed':'2021-07-27T01:11:05+07:00', 'timestamp':'1595787065171'}]
        post_filter = service.filtering(data)
        result = service.aggregating(post_filter)
        for res in result['result']:
            if res['end_week'] == '2021-07-04':
                if res['min_size'] != 70:
                    errors.append(f'Wrong value in min_size field for week {res["end_week"]}')
                if res['max_size'] != 70:
                    errors.append(f'Wrong value in max_size field for week {res["end_week"]}')
                if res['med_size'] != 70:
                    errors.append(f'Wrong value in med_size field for week {res["end_week"]}')
                if res['mean_size'] != 70:
                    errors.append(f'Wrong value in mean_size field for week {res["end_week"]}')
                if res['min_price'] != 30000:
                    errors.append(f'Wrong value in min_price field for week {res["end_week"]}')
                if res['max_price'] != 30000:
                    errors.append(f'Wrong value in max_price field for week {res["end_week"]}')
                if res['med_price'] != 30000:
                    errors.append(f'Wrong value in med_price field for week {res["end_week"]}')
                if res['mean_price'] != 30000:
                    errors.append(f'Wrong value in mean_price field for week {res["end_week"]}')

            elif res['end_week'] == '2021-08-01':
                if res['min_size'] != 50:
                    errors.append(f'Wrong value in min_size field for week {res["end_week"]}')
                if res['max_size'] != 120:
                    errors.append(f'Wrong value in max_size field for week {res["end_week"]}')
                if res['med_size'] != 85:
                    errors.append(f'Wrong value in med_size field for week {res["end_week"]}')
                if res['mean_size'] != 85:
                    errors.append(f'Wrong value in mean_size field for week {res["end_week"]}')
                if res['min_price'] != 20000:
                    errors.append(f'Wrong value in min_price field for week {res["end_week"]}')
                if res['max_price'] != 45000:
                    errors.append(f'Wrong value in max_price field for week {res["end_week"]}')
                if res['med_price'] != 32500:
                    errors.append(f'Wrong value in med_price field for week {res["end_week"]}')
                if res['mean_price'] != 32500:
                    errors.append(f'Wrong value in mean_price field for week {res["end_week"]}')

        assert not errors, ','.join(errors)