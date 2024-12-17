function zdy_dm5_jm(jm_str, t)
	local salt = "v50gjcy"
	local combined = jm_str .. salt .. t
	local md5_val = utils.md5(combined)
	return md5_val
end

function custom_head(vurl)
  local headers = {}
  headers["User-Agent"] = ""
  headers["Referer"] = ""
  if type(vurl) == "string" and string.find(vurl, "aliyuncs.com") then
    headers["Referer"] = "https://survey.goofish.com"
  end
  return headers
end

function parser(source)
  local parsers = { "http://yhhy.xj.zshtys888.com/vo1v03.php?url=" }
  local timestamp = utils.timestamp()
  local modified_source = source .. "&t=" .. timestamp

  for i = 1, #parsers do
    local get_header = {}
    get_header["x-time"] = timestamp
    get_header["x-form"] = device_info["platform"]
    get_header["x-sign1"] = zdy_dm5_jm(device_info["app_version"], timestamp)
    get_header["x-sign2"] = zdy_dm5_jm(source, timestamp)
    local response_body = {}
    local full_url = parsers[i] .. modified_source
    local jsonData = httpGet(full_url, {header = get_header})
    local obj, pos, err = json.decode(jsonData, 1, nil)
    if err then
      local jsonDataDenc = utils.aes128cbc_decrypt("wcyjmnnnawozmydn", "wcivwyjmlnzbhlmq", jsonData)
      obj, pos, err = json.decode(jsonDataDenc, 1, nil)
    end
    if not err and (obj.code == 200 or obj.code == "200") then
      local headers = custom_head(obj.url)

      if obj.type ~= nil and (obj.type == "multi" or obj.type == "mp4" or obj.type == "m3u8" or obj.type == "hls" or obj.type == "flv") then
        if device_info ~= nil then
          if device_info["platform"] == "Android" then
            -- Android
          elseif device_info["platform"] == "Mac" then
            -- Mac
          elseif device_info["platform"] == "Windows" then
            -- Windows
          elseif device_info["platform"] == "TV" then
            -- TV
          else
            -- 其他
          end
        end

        local url = obj.url
        local header = ""

        if obj.type == "multi" then
          url = json.encode(url)
        end

        if next(headers) ~= nil then
          header = json.encode(headers)
        end

        return "OK", url, header, obj.type
      end
    end
    response_body = nil
  end
  return "ERROR", "", "", ""
end