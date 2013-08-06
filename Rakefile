require "safe_yaml"

class Hash
  # Merges self with another hash, recursively.
  #
  # This code was lovingly stolen from some random gem:
  # http://gemjack.com/gems/tartan-0.1.1/classes/Hash.html
  #
  # Thanks to whoever made it.
  def deep_merge(hash)
    target = dup
    hash.keys.each do |key|
      if hash[key].is_a? Hash and self[key].is_a? Hash
        target[key] = target[key].deep_merge(hash[key])
        next
      end
      target[key] = hash[key]
    end
    target
  end
end

def wait_and_kill(pids)
  trap("INT") {
    pids.each { |pid| Process.kill(3, pid) rescue Errno::ESRCH }
    exit 0
  }

  pids.each { |pid| Process.wait(pid) }
end

def data_files
  Dir["_data/*"]
end

def read_and_merge(files)
  contents = []
  files.each do |file|
    contents << File.read(file)
  end
  contents.join("\n")
end

desc "Compile site"
task :compile do
  jekyllPid = Process.spawn("jekyll build")
  compassPid = Process.spawn("compass compile ./_sass -e #{ENV['ENV'] || 'development'}")
  wait_and_kill [jekyllPid, compassPid]
end

desc "Watch"
task :watch do
  jekyllPid = Process.spawn("jekyll build --watch")
  compassPid = Process.spawn("compass watch ./_sass -e #{ENV['ENV'] || 'development'}")
  wait_and_kill [jekyllPid, compassPid]
end

desc "Preview"
task :preview do
  jekyllPid = Process.spawn("jekyll serve --watch --config #{data_files.join(",")}")
  compassPid = Process.spawn("compass watch ./_sass -e #{ENV['ENV'] || 'development'}")
  wait_and_kill [jekyllPid, compassPid]
end

desc "Deploy to GH Pages"
task :deploy do
  File.open('_config.yml', 'wb') do |f|
    f.write(read_and_merge(data_files))
  end
  sh "compass compile ./_sass -e production"
  if `git status --porcelain 2> /dev/null`.strip.include?("_config.yml")
    sh "git add _config.yml"
    sh "git commit -qm 'Update _config.yml based on contents of _data directory"
  end
  sh "git push"
end
